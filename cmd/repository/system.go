package repository

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/tech-thinker/gozen/utils"
	"github.com/tech-thinker/gozen/wrappers"
)

type SystemRepo interface {
	Write(templatePath string, outputPath string, data interface{}) error
	ExecShell(command string, args ...string) ([]string, error)
	ExecShellRaw(command string, args ...string) ([]byte, error)
}

type systemRepo struct {
	templatesFS    embed.FS
	shellRepo      wrappers.ShellWrapper
	fileSystemRepo wrappers.FileSystemWrapper
}

// Write: generate code and write to file
func (helper *systemRepo) Write(templatePath string, outputPath string, data interface{}) error {
	baseDir := filepath.Dir(outputPath)
	err := helper.fileSystemRepo.CreateDirectory(baseDir)
	if err != nil {
		return err
	}
	tpl, err := utils.GenerateCode(helper.templatesFS, templatePath, data)
	if err != nil {
		return err
	}
	tpl = utils.ApplyEscapeChar(tpl)
	return helper.fileSystemRepo.WriteFile(outputPath, tpl)
}

// ExecShell: execute shell command and return output as string slice
func (helper *systemRepo) ExecShell(command string, args ...string) ([]string, error) {
	fmt.Printf(`%s %+v\n`, command, args)
	output, err := helper.shellRepo.Exec(command, args...)
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	return lines, nil
}

// ExecShellRaw: execute shell command and return output as byte array
func (helper *systemRepo) ExecShellRaw(command string, args ...string) ([]byte, error) {
	fmt.Printf(`%s %+v\n`, command, args)
	output, err := helper.shellRepo.Exec(command, args...)
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}
	return output, nil
}

// NewSystemRepo returns a new SystemRepo
func NewSystemRepo(
	tpl embed.FS,
	shellRepo wrappers.ShellWrapper,
	fileSystemRepo wrappers.FileSystemWrapper,
) SystemRepo {
	return &systemRepo{
		templatesFS:    tpl,
		shellRepo:      shellRepo,
		fileSystemRepo: fileSystemRepo,
	}
}
