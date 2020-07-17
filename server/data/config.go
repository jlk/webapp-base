package data // import "github.com/jlk/webapp-base/server/data"

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config - global config for the app
var Config = viper.New()

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("j")

	Config.SetConfigName("webserver")
	Config.AddConfigPath("/etc/webserver/")
	Config.AddConfigPath(".")
	if err := Config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			logrus.Fatalf("Error while reading configuration file: %s\n", err.Error())
		}
	}

	Config.SetDefault("ListenPort", 4000)
	Config.SetDefault("dbSSLMode", "disable")
	Config.SetDefault("dbHost", "localhost")
}
