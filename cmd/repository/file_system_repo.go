package repository

import (
	"fmt"
	"os"
)

// FileSystemRepo is an abstructtion of file handling system library, no test cases required for this
type FileSystemRepo interface {
	CreateDirectory(path string) error
	WriteFile(path, data string) error
}

type fileSystemRepo struct {
}

func (repo *fileSystemRepo) CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (repo *fileSystemRepo) WriteFile(path, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, `%s`, data)
	if err != nil {
		return err
	}
	return nil
}

func NewFileSystemRepo() FileSystemRepo {
	return &fileSystemRepo{}
}
