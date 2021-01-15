package config

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func loadUATSettings() {
	secretPath := "/mnt/secrets-store"

	viper.AddConfigPath(secretPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if mqusername, err := ioutil.ReadFile(fmt.Sprintf("%s/mqusername-uat", secretPath)); err == nil {
		viper.SetDefault("mqusername", string(mqusername))
	}

	if mqpassword, err := ioutil.ReadFile(fmt.Sprintf("%s/mqpassword-uat", secretPath)); err == nil {
		viper.SetDefault("mqpassword", string(mqpassword))
	}

	if mqhost, err := ioutil.ReadFile(fmt.Sprintf("%s/mqhost-uat", secretPath)); err == nil {
		viper.SetDefault("mqhost", string(mqhost))
	}

	viper.SetDefault("mqport", 5672)
}
