package config_dto

// AppConfig is the struct that holds the configuration values for the application
type AppConfig struct {
	SettingPath string `yaml:"setting_path"`
	Env         string `yaml:"env"`
	Prefix      string `yaml:"prefix"`
}
