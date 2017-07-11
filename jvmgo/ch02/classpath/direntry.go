package classpath

import (
	"path/filepath"
	"io/ioutil"
)

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
