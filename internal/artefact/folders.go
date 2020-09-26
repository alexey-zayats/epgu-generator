package artefact

import (
	"github.com/pkg/errors"
	"os"
	"path"
)

// Folders ...
type Folders struct {
	rootDir string
	data    map[string]string
}

// NewFolders ...
func NewFolders(root string, fromCode string) *Folders {

	data := map[string]string{
		"forms/svcspec":          path.Join(root, fromCode, "svcspec"),
		"forms/svcspec/rollback": path.Join(root, fromCode, "svcspec", "rollback"),
	}

	return &Folders{
		rootDir: root,
		data:    data,
	}
}

// Struct ...
func (f *Folders) Struct() map[string]string {
	return f.data
}

// MakeStruct ...
func (f *Folders) MakeStruct() error {

	if err := os.MkdirAll(f.data["forms/svcspec/rollback"], os.ModePerm); err != nil {
		return errors.Wrapf(err, "unable call os.MkdirAll(%s)", f.data["forms/svcspec/rollback"])
	}

	return nil
}
