package models

import "{{.PackageName}}/instance/registry"

type Health struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	HeartBeat int64  `json:"heart_beat"`
}

func init() {
	registry.RegisterModel(new(Health))
}
