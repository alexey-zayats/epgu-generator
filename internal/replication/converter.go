package replication

import (
	"context"
	"epgu-generator/internal/artefact"
	"epgu-generator/internal/config"
	"epgu-generator/internal/model"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"os"
	"path"
)

// Converter ...
type Converter struct {
	config  *config.Config
	content *artefact.Content
}

// DI ...
type DI struct {
	dig.In
	Config  *config.Config
	Content *artefact.Content
}

// NewConverter ...
func NewConverter(di DI) *Converter {
	return &Converter{
		config:  di.Config,
		content: di.Content,
	}
}

// Convert ...
func (c *Converter) Convert(ctx context.Context, reg *model.Replication) error {

	recordPath := path.Join(c.config.Dir.Tmp, reg.FormCode)

	for _, reg := range reg.Items {
		folders := artefact.NewFolders(recordPath, fmt.Sprintf("form.61.%s", reg.DepartmentCode))

		// Создаем структуру папок архива
		if err := folders.MakeStruct(); err != nil {
			return errors.Wrap(err, "unable make dir struct")
		}

		// Генерируем скрипты по шаблонам
		if err := c.content.Prepare(reg, folders.Struct()); err != nil {
			return errors.Wrapf(err, "unable to prepare content for %s", reg.ServiceTargetID)
		}
	}

	// Архивируем артефакты
	recordArchiver := artefact.NewArchiver(recordPath, c.config.Dir.Artefact, reg.FormCode+".zip")
	if err := recordArchiver.Compose(); err != nil {
		return errors.Wrapf(err, "unable to composer archive for %s", reg.FormCode)
	}

	// Очищаем
	if err := os.RemoveAll(recordPath); err != nil {
		return errors.Wrapf(err, "unable to call os.RemoveAll(%s)", recordPath)
	}

	return nil
}
