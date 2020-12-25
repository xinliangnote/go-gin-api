package configs

import (
	"github.com/xinliangnote/go-gin-api/pkg/env"

	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	DB struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"db"`

	Signature struct {
		Secrets map[string]string `mapstructure:"secrets"`
	} `mapstructure:"signature"`
}

func init() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}

func ProjectName() string {
	return "go-gin-api"
}
