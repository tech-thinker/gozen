package config

import "github.com/spf13/viper"

type Configuration interface {
	Version() string
	AppName() string
	APIPort() string
    
    DbDriver() string
	DbHost() string
	DbPort() string
	DbUser() string
	DbPass() string
	DbName() string
}
type configuration struct {
	version string
	appName string
	apiPort string

    db_driver string
    db_host string
    db_port string
    db_user string
    db_pass string
    db_name string
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

// DbDriver returns database driver
func (config *configuration) DbDriver() string {
    return config.db_driver
}

// DbHost returns database host
func (config *configuration) DbHost() string {
    return config.db_host
}

// DbPort returns database port
func (config *configuration) DbPort() string {
    return config.db_port
}

// DbUser returns database user
func (config *configuration) DbUser() string {
    return config.db_user
}

// DbPass returns database password
func (config *configuration) DbPass() string {
    return config.db_pass
}

// DbName returns database name
func (config *configuration) DbName() string {
    return config.db_name
}

func Init(
	env *viper.Viper,
) Configuration {

	config := &configuration{}

	env.AutomaticEnv()
	config.version = env.GetString("version")
	config.apiPort = env.GetString("app_name")
	config.apiPort = env.GetString("api_port")
    config.db_driver = env.GetString("db_driver")
    config.db_host = env.GetString("db_host")
    config.db_port = env.GetString("db_port")
    config.db_user = env.GetString("db_user")
    config.db_pass = env.GetString("db_pass")
    config.db_name = env.GetString("db_name")
	return config
}
