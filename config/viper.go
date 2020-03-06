package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func InitializeConfiguration(prefix, serviceName, configFile, configPath string) {
	log.Println("[Warning] InitializeConfiguration is deprecated. Use LoadRemoteConfiguration")
	//viper.SetEnvPrefix(prefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetConfigName(configFile)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(fmt.Sprintf("/etc/%s", serviceName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", serviceName))
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s", err)
	}
}
