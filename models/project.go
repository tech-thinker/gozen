package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tech-thinker/gozen/constants"
)

type Project struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
	Driver      string `json:"driver"`
	WorkingDir  string `json:"-"`
}

func (p Project) Validate() error {
	if len(p.PackageName) == 0 {
		return errors.New("Please provide package name.")
	}
	return nil
}

func (p *Project) AutoFixes() {
	if len(p.AppName) == 0 {
		p.AppName = p.PackageName[strings.LastIndex(p.PackageName, "/")+1:]
	}
}

func (p Project) ToJSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

func (p *Project) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), p)
}

func (p *Project) LoadFromJsonFile() error {
	filename := fmt.Sprintf(`%s/gozen.json`, constants.CURRENT_WORKING_DIRECTORY)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	return p.FromJSON(string(bytes))
}

func (p Project) WriteToJsonFile() error {
	workDir := constants.CURRENT_WORKING_DIRECTORY
	if len(p.WorkingDir) > 1 {
		workDir = p.WorkingDir
	}

	filename := fmt.Sprintf(`%s/%s/gozen.json`, workDir, p.AppName)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(p.ToJSON())
	if err != nil {
		return err
	}
	return nil
}
