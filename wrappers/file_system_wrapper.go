package wrappers

import (
	"fmt"
	"os"
)

// FileSystemWrapper is an abstructtion of file handling system library, no test cases required for this
type FileSystemWrapper interface {
	CreateDirectory(path string) error
	Destroy(path string) error
	WriteFile(path, data string) error
}

type fileSystemWrapper struct {
}

func (repo *fileSystemWrapper) CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (repo *fileSystemWrapper) Destroy(path string) error {
	return os.Remove(path)
}

func (repo *fileSystemWrapper) WriteFile(path, data string) error {
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

func NewFileSystemWrapper() FileSystemWrapper {
	return &fileSystemWrapper{}
}
