package data // import "github.com/jlk/webapp-base/server/data"

import "github.com/spf13/viper"

// Config - global config for the app
var Config = viper.New()

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("j")

	// Move these somewhere else in time
	Config.SetDefault("ListenPort", 3000)
	// mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	// or https://godoc.org/github.com/globalsign/mgo#Dial
	Config.SetDefault("mgoURL", "mongodb://admin:mongocloud@localhost:27017/admin")
	// TODO: mgo tls support
}
