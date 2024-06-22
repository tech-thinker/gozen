package models

import "{{.PackageName}}/instance/registry"

type Health struct {
    Success bool `json:"success"`
    Message string `json:"message"`
}

func init() {
    registry.RegisterModel(new(Health))
}
