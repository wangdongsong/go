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