package instance

import (
	"{{.PackageName}}/config"
)

type Instance interface {
	Destroy() error
}

type instance struct {
	// TODO: define singleton object here
}

// Destroy closes the connections & cleans up the instance
func (instance *instance) Destroy() error {
	return nil
}


// Init initializes the instance
func Init(cfg config.Configuration) Instance {
	instance := &instance{}
    // TODO: initialize singleton object here

	return instance
}