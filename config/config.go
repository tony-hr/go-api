package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(params ...string) *viper.Viper {
	env := "development"
	if len(params) > 0 {
		env = params[0]
	}

	viper.SetConfigType("json")
	viper.AddConfigPath("config")
	viper.SetConfigName(env)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	return viper.GetViper()
}
