package artefact

import (
	"epgu-generator/internal/config"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"io"
	"path"
	"text/template"
)

// Template ...
type Template struct {
	path string
	dict map[string]*template.Template
}

// TemplateDI ...
type TemplateDI struct {
	dig.In
	Config *config.Config
}

// NewTemplates ...
func NewTemplates(di TemplateDI) (*Template, error) {

	t := &Template{
		path: di.Config.Dir.Template,
		dict: make(map[string]*template.Template),
	}
	return t, nil
}

// AddFile ...
func (t *Template) AddFile(item string) error {

	filePath := path.Join(t.path, item)

	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return errors.Wrapf(err, "unable to parse template file %s", filePath)
	}

	t.dict[item] = tmpl

	return nil
}

// Execute ...
func (t *Template) Execute(wr io.Writer, name string, data interface{}) error {

	tmpl, ok := t.dict[name]
	if ok == false {
		return errors.New(fmt.Sprintf("unable to find template %s", name))
	}

	if err := tmpl.Execute(wr, data); err != nil {
		return errors.Wrapf(err, "unable to call template.Execute() for %s", name)
	}

	return nil
}
