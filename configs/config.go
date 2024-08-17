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
	var cfg conf
	viper.AddConfigPath(path)
	viper.SetConfigName(".env" + env)
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			port := viper.Get("PORT")
			dbConn := viper.Get("DB_CONNECTION")
			if dbConn == nil || port == nil {
				log.Fatal("DB_CONNECTION or PORT not found during configuration")
			}
			cfg.DBConn = dbConn.(string)
			cfg.Port = port.(string)
			return &cfg
		} else {
			log.Fatal(err)
		}
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
