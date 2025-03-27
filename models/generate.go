package models

type Generator struct {
	Project
	Type          string `json:"type"`
	InterfaceName string `json:"interfaceName"`
	StructName    string `json:"structName"`
	RouteName     string `json:"routeName"`
}
