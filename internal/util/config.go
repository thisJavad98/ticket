package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APPPORT      string `mapstructure:"APP_PORT"`
	APPNAME      string `mapstructure:"APP_NAME"`
	APPDEBUG     string `mapstructure:"APP_DEBUG" `
	DBCONNECTION string `mapstructure:"DB_CONNECTION"`
	DBHOST       string `mapstructure:"DB_HOST"`
	DBPORT       string `mapstructure:"DB_PORT"`
	DBDATABASE   string `mapstructure:"DB_DATABASE"`
	DBUSERNAME   string `mapstructure:"DB_USERNAME"`
	DBPASSWORD   string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("cannot configoration app!!")
		return
	}

	err = viper.Unmarshal(&config)
	fmt.Println("config app successfully!")
	return
}
