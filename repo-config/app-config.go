package repo_config

import (
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"flag"
	"fmt"
	"os"
)

func NewApplication() (*config_dto.AppConfig, error) {
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	env := cmdLine.String("env", "qa", "Environment(dev, qa, stg, prod)")

	err := cmdLine.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("NewApplication: cmdLine.Parse ERROR=%v", err)
		return nil, err
	}
	appConfig := config_dto.AppConfig{Prefix: "/api", SettingPath: "repo-config/settings/settings.%s.yml", Env: *env}
	return &appConfig, nil
}
