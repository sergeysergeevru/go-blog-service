package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Cfg struct {
	Server Server `mapstructure:"server"`
}

type Server struct {
	Port   int    `mapstructure:"port"`
	Prefix string `mapstructure:"prefix"`
}

func ReadConfig(filePath string) (*Cfg, error) {
	configReader := viper.New()
	configReader.AllowEmptyEnv(true)
	configReader.AutomaticEnv()
	configReader.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	configReader.SetConfigFile(filePath)
	if err := configReader.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf Cfg
	if err := configReader.Unmarshal(&conf); err != nil {
		return nil, err

	}

	return &conf, nil
}
