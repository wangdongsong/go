package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

type ZipEntry struct{
	absPath string
}

func newZipEntry(path string) *ZipEntry{
	absPath, error := filepath.Abs(path)

	if error != nil{
		panic(error)
	}

	return &ZipEntry{absPath}
}

func(self *ZipEntry) String() string{
	return self.absPath
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error)  {
	r, error := zip.OpenReader(self.absPath)

	if error != nil {
		return nil, nil, error
	}

	defer r.Close()

	for _, f := range r.File{
		if f.Name == className {
			rc, error := f.Open()
			if error != nil {
				return nil, nil, error
			}
			defer rc.Close()

			data, error := ioutil.ReadAll(rc)
			if error != nil {
				return nil, nil, error
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found:" + className)
}