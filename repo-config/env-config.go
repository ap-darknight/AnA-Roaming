package repo_config

import (
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"fmt"
	"github.com/jinzhu/configor"
)

func NewConfig(appConfig *config_dto.AppConfig) (*config_dto.Config, error) {
	filename := fmt.Sprintf(appConfig.SettingPath, appConfig.Env)
	config := &config_dto.Config{}
	if err := configor.Load(config, filename); err != nil {
		return nil, err
	}

	config.AppConfig = *appConfig
	return config, nil
}
