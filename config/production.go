package config

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func loadProdSettings() {
	secretPath := "/mnt/secrets-store"

	viper.AddConfigPath(secretPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if mqusername, err := ioutil.ReadFile(fmt.Sprintf("%s/mqusername-prod", secretPath)); err == nil {
		viper.SetDefault("mqusername", string(mqusername))
	}

	if mqpassword, err := ioutil.ReadFile(fmt.Sprintf("%s/mqpassword-prod", secretPath)); err == nil {
		viper.SetDefault("mqpassword", string(mqpassword))
	}

	if mqhost, err := ioutil.ReadFile(fmt.Sprintf("%s/mqhost-prod", secretPath)); err == nil {
		viper.SetDefault("mqpassword", string(mqhost))
	}

	viper.SetDefault("mqport", 5672)
}
