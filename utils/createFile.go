package utils

import (
	"os"
	"path/filepath"
)

func CreateFile(fileName string, absPath string) *os.File {
	composePath, _ := filepath.Abs(absPath + "/" + fileName)
	file, err := os.Create(composePath)
	Check(err)
	defer file.Close()

	return file
}
