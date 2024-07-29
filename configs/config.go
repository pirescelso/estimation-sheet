package configs

import (
	"log"

	"github.com/spf13/viper"
)

type conf struct {
	DBConn string `mapstructure:"DB_CONNECTION"`
	Port   string `mapstructure:"PORT"`
}

func LoadConfig(path string, env string) *conf {
	var cfg *conf
	viper.AddConfigPath(path)
	viper.SetConfigName(".env" + env)
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
