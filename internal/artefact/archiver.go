package artefact

import (
	"archive/zip"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

// Archiver ...
type Archiver struct {
	src  string
	dst  string
	name string
}

// NewArchiver ...
func NewArchiver(src, dst, name string) *Archiver {
	return &Archiver{
		src:  src,
		dst:  dst,
		name: name,
	}
}

// Compose ...
func (a *Archiver) Compose() error {

	// Get a Buffer to Write To
	archiveFile := path.Join(a.dst, a.name)
	outFile, err := os.Create(archiveFile)
	if err != nil {
		return errors.Wrapf(err, "unable to call os.Create(%s)", archiveFile)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			logrus.WithError(err).Error("unable to close file %s", archiveFile)
		}
	}()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	if err := a.addFiles(w, a.src+"/", ""); err != nil {
		return errors.Wrapf(err, "unable to add files from(%s) to the archive", a.src)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return errors.Wrap(err, "unable to close zip.Writer")
	}
	return nil
}

func (a *Archiver) addFiles(w *zip.Writer, basePath, baseInZip string) error {

	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return errors.Wrapf(err, "unable to call ioutil.ReadDir(%s)", basePath)
	}

	for _, file := range files {

		if !file.IsDir() {

			filePath := basePath + file.Name()
			dat, err := ioutil.ReadFile(filePath)
			if err != nil {
				return errors.Wrapf(err, "unable to call ioutil.ReadFile(%s)", filePath)
			}

			// Add some files to the archive.
			zipPath := baseInZip + file.Name()
			f, err := w.Create(zipPath)
			if err != nil {
				return errors.Wrapf(err, "unable to add file to the archive", zipPath)
			}
			_, err = f.Write(dat)
			if err != nil {
				return errors.Wrapf(err, "unable to Write file content ", zipPath)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			newBaseInZip := baseInZip + file.Name() + "/"
			if err := a.addFiles(w, newBase, newBaseInZip); err != nil {
				return errors.Wrapf(err, "unable to add files from(%s) to archive(%s)", newBase, newBaseInZip)
			}
		}
	}

	return nil
}
