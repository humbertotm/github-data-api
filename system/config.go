package system

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Cfg EnvConfig

type EnvConfig struct {
	Mode    string `envconfig:"mode"`
	LogFile string `envconfig:"log_file"`
	DbURL   string `envconfig:"db_url"`
}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Cfg); err != nil {
		return err
	}

	return nil
}

func IsDev() bool {
	return Cfg.Mode == "dev"
}
