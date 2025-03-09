package util

import (
	"os"
	"path/filepath"
)

type FileData struct {
	IsDir bool
	Path  string
}

func ScanFiles(dir string) ([]FileData, error) {
	var fileData []FileData

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {

			file := FileData{
				IsDir: info.IsDir(),
				Path:  path,
			}

			fileData = append(fileData, file)
		}
		return nil
	})
	if err != nil {
		return fileData, err
	}
	return fileData, nil
}
