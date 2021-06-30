package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     int      `mapstructure:"port"`
	Jwt      string   `mapstructure:"jwt"`
	Key      string   `mapstructure:"key"`
	Database Database `mapstructure:"database"`
}

type Database struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

var Conf = new(Config)

func SetUp() {
	viper.SetConfigFile("./config/config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Unmarshal conf failed, err: %s", err)
	}
}
