package classpath

import (
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
)

const pathlistSeqarator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry{
	if strings.Contains(path, pathlistSeqarator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP")) {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

type DirEntry struct{
	absDir string
}

func newDirEntry(path string) *DirEntry{
	absDir, error := filepath.Abs(path)
	if error != nil{
		//终止函数执行
		panic(error)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error)  {
	fileName := filepath.Join(self.absDir, className)
	data, error := ioutil.ReadFile(fileName)
	return data, self, error
}

func (self *DirEntry) String() string  {
	return self.absDir
}