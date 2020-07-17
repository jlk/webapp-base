package data // import "github.com/jlk/webapp-base/server/data"

import "github.com/spf13/viper"

// Config - global config for the app
var Config = viper.New()

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("j")

	// Move these somewhere else in time
	Config.SetDefault("ListenPort", 4000)
	Config.SetDefault("dbSSLMode", false)
	Config.SetDefault("dbHost", "mongodb://admin:mongocloud@localhost:27017/admin")
}
