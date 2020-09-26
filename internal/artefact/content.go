package artefact

import (
	"epgu-generator/internal/config"
	"epgu-generator/internal/model"
	"epgu-generator/internal/util"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
	"path"
	"strings"
)

// ItemKind ...
type ItemKind int

const (
	// ItemNone ...
	ItemNone ItemKind = iota
	// ItemTemplate ...
	ItemTemplate
	// ItemCopy ...
	ItemCopy
)

// Item ...
type Item struct {
	Path string
	Kind ItemKind
}

// Content ...
type Content struct {
	items        []Item
	templatePath string
	template     *Template
	incrementer  *Incrementer
}

// ContentDI ...
type ContentDI struct {
	dig.In
	Config      *config.Config
	Template    *Template
	Incrementer *Incrementer
}

// NewContent ...
func NewContent(di ContentDI) (*Content, error) {

	items := []Item{
		{"forms/svcspec/inc-pguapi-pub-form_target.sql", ItemTemplate},
		{"forms/svcspec/rollback/inc-pguapi-pub-form_target.sql", ItemTemplate},
	}

	templatePath := di.Config.Dir.Template

	for _, item := range items {
		if item.Kind != ItemTemplate {
			continue
		}
		if err := di.Template.AddFile(item.Path); err != nil {
			return nil, errors.Wrapf(err, "unable add template file %s", item.Path)
		}
	}

	return &Content{
		template:     di.Template,
		items:        items,
		templatePath: templatePath,
		incrementer:  di.Incrementer,
	}, nil
}

// Prepare ...
func (c *Content) Prepare(reg *model.Registry, folders map[string]string) error {

	applicant := make([]string, len(reg.ApplicantType))
	for i, item := range reg.ApplicantType {
		switch item {
		case model.ApplicantIE:
			applicant[i] = "SOLE_PROPRIETOR"
		case model.ApplicantNP:
			applicant[i] = "PERSON"
		case model.ApplicantLE:
			applicant[i] = "LEGAL"
		}
	}

	var useSignature string
	switch reg.UseSignature {
	case true:
		useSignature = "EDS_MANDATORY"
	case false:
		useSignature = "EDS_NOT_SUPPORTED"
	}

	meta := Meta{
		DepartmentName:    reg.DepartmentName,
		DepartmentCode:    reg.DepartmentCode,
		ServiceName:       reg.ServiceName,
		ServiceTargetName: reg.ServiceTargetName,
		ServiceTargetID:   reg.ServiceTargetID,
		ServiceFormCode:   reg.ServiceFormCode,
		ApplicantType:     strings.Join(applicant, ","),
		Signature:         useSignature,
		UnlinkService:     reg.UnlinkService,
		Change:            reg.Change,
	}

	inc := fmt.Sprintf("%03d", c.incrementer.Get(reg.DepartmentCode + reg.ServiceFormCode))

	for _, item := range c.items {

		switch item.Kind {
		case ItemCopy:

			paths := strings.Split(item.Path, "/")
			name := paths[len(paths)-1]
			dir := strings.Join(paths[:len(paths)-1], "/")

			src := path.Join(c.templatePath, item.Path)
			dst := path.Join(folders[dir], name)

			if err := util.CopyFile(src, dst); err != nil {
				return errors.Wrapf(err, "unable to copy file from %s to %s", src, dst)
			}

		case ItemTemplate:

			paths := strings.Split(item.Path, "/")
			name := paths[len(paths)-1]
			dir := strings.Join(paths[:len(paths)-1], "/")

			fileName := strings.ReplaceAll(name, "form", reg.ServiceFormCode)
			fileName = strings.ReplaceAll(fileName, "target", reg.ServiceTargetID)
			fileName = strings.ReplaceAll(fileName, "inc", inc)

			filePath := path.Join(folders[dir], fileName)

			file, err := os.Create(filePath)
			if err != nil {
				return errors.Wrapf(err, "unable to create file %s", filePath)
			}

			writer := transform.NewWriter(file, charmap.Windows1251.NewEncoder())

			if err := c.template.Execute(writer, item.Path, meta); err != nil {
				_ = file.Close()
				return errors.Wrapf(err, "unable to call template.ExecuteTemplate(%s)", name)
			}

			if err := file.Close(); err != nil {
				logrus.WithError(err).Errorf("unable close file %s", filePath)
			}
		}
	}

	return nil
}
