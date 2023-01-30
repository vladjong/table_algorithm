package flag

import (
	"fmt"
	"os"
	"path"
)

func GetFilename() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("[getFilename]:empty filename")
	}
	filename := os.Args[1]
	ext := path.Ext(filename)
	if ext != ".csv" {
		return "", fmt.Errorf("[getFilename]:incorrect extension file:%v, need:.csv", ext)
	}
	return filename, nil
}
