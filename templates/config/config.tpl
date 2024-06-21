package config

import "github.com/spf13/viper"

type Configuration interface {
	Version() string
	AppName() string
	APIPort() string
}
type configuration struct {
	version string
	appName string
	apiPort string
}

// Version returns version
func (cfg *configuration) Version() string {
	return cfg.version
}

// AppName returns AppName
func (cfg *configuration) AppName() string {
	return cfg.appName
}

// APIPort returns port
func (cfg *configuration) APIPort() string {
	return cfg.apiPort
}

func Init(
	env *viper.Viper,
) Configuration {

	config := &configuration{}

	env.AutomaticEnv()
	config.version = env.GetString("version")
	config.apiPort = env.GetString("app_name")
	config.apiPort = env.GetString("api_port")
	return config
}
