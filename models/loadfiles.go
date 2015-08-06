package models

import (
	"os"
)

func Listfiles(dir string) []string {
	f, err := os.Open(dir)
	if nil != err {
		return nil
	}
	dirs, err := f.Readdir(-1)
	if nil != err {
		return nil
	}
	ret := make([]string, 0, len(dirs))
	for _, it := range dirs {
		if !it.IsDir() {
			ret = append(ret, it.Name())
		}
	}
	return ret
}
