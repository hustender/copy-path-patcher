package fileops

import (
	"os"
	"path/filepath"
)

func GetSubFiles(path string) ([]os.FileInfo, error) {
	var files []os.FileInfo
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
