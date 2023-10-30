package configs

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error while attempting get key [%s]: %e", key, err)
	}

	return viper.GetString(key)
}
