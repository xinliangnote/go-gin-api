package configs

import (
	"github.com/xinliangnote/go-gin-api/pkg/env"

	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	Mail struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		To   string `mapstructure:"to"`
	} `mapstructure:"mail"`

	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`

	Aes struct {
		Key string `mapstructure:"key"`
		Iv  string `mapstructure:"iv"`
	} `mapstructure:"aes"`

	Rsa struct {
		Private string `mapstructure:"private"`
		Public  string `mapstructure:"public"`
	} `mapstructure:"rsa"`
}

func init() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("toml")
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
