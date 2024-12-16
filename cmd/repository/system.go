package repository

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/utils"
	"github.com/tech-thinker/gozen/wrappers"
)

// SystemRepo is an interface that defines methods for interacting with the system.  This includes writing files, executing shell commands, and managing the file system.
type SystemRepo interface {
	// Write generates code from a template and writes it to a file.
	Write(appDir string, templatePath string, destination string, data interface{}) error
	// WriteAll generates code from multiple templates and writes them to their respective files.
	WriteAll(appDir string, fileConfigs []models.FileConfig, data interface{}) error
	// ExecShell executes a shell command with arguments and returns the output as a string slice.
	ExecShell(command string, args ...string) ([]string, error)
	// ExecShellRaw executes a shell command with arguments and returns the raw byte output.
	ExecShellRaw(command string, args ...string) ([]byte, error)
}

type systemRepo struct {
	templatesFS    embed.FS
	shellRepo      wrappers.ShellWrapper
	fileSystemRepo wrappers.FileSystemWrapper
}

// Write: generate code and write to file
func (repo *systemRepo) Write(appDir string, templatePath string, destination string, data interface{}) error {
	outputPath := filepath.Join(appDir, destination)
	baseDir := filepath.Dir(outputPath)
	err := repo.fileSystemRepo.CreateDirectory(baseDir)
	if err != nil {
		return err
	}
	tpl, err := utils.GenerateCode(repo.templatesFS, templatePath, data)
	if err != nil {
		return err
	}
	tpl = utils.ApplyEscapeChar(tpl)
	return repo.fileSystemRepo.WriteFile(outputPath, tpl)
}

// WriteAll: generate codes and write to files
func (repo *systemRepo) WriteAll(appDir string, fileConfigs []models.FileConfig, data interface{}) error {
	var errs []error
	for _, fileConfig := range fileConfigs {
		err := repo.Write(appDir, fileConfig.TemplatePath, fileConfig.Destination, data)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("WriteAll encountered errors: %v", errs)
	}
	return nil
}

// ExecShell: execute shell command and return output as string slice
func (repo *systemRepo) ExecShell(command string, args ...string) ([]string, error) {
	fmt.Printf(`%s %+v\n`, command, args)
	output, err := repo.shellRepo.Exec(command, args...)
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	return lines, nil
}

// ExecShellRaw: execute shell command and return output as byte array
func (repo *systemRepo) ExecShellRaw(command string, args ...string) ([]byte, error) {
	fmt.Printf(`%s %+v\n`, command, args)
	output, err := repo.shellRepo.Exec(command, args...)
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
