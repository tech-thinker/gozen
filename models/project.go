package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tech-thinker/gozen/constants"
)

type Project struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
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
    filename := fmt.Sprintf(`%s/%s/gozen.json`, constants.CURRENT_WORKING_DIRECTORY, p.AppName)
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
