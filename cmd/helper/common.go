package helper

import (
	"embed"
	"path/filepath"

	"github.com/tech-thinker/gozen/utils"
)

type CommonHelper interface {
    Write(templatePath string, outputPath string, data interface{}) error
}

type commonHelper struct {
    templatesFS embed.FS
}

// Write: generate code and write to file
func (helper *commonHelper) Write(templatePath string, outputPath string, data interface{}) error {
    baseDir := filepath.Dir(outputPath)
    err := utils.CreateDirectory(baseDir)
    if err != nil {
        return err
    }
    tpl, err := utils.GenerateCode(helper.templatesFS, templatePath, data)
    if err != nil {
        return err
    }
    return utils.WriteFile(outputPath, tpl)
}

// NewCommonHelper returns a new CommonHelper
func NewCommonHelper(tpl embed.FS) CommonHelper {
    return &commonHelper{
        templatesFS: tpl,
    }
}
