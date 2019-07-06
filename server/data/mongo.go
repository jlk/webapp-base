package data // import "github.com/jlk/webapp-base/server/data"

import (
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

// MongoClient - Our mongo db session
var MongoClient *mgo.Session

func init() {
	var err error

	mgoURL := Config.GetString("mgoURL")
	if len(mgoURL) == 0 {
		logrus.Fatal("No value set for configuration variable mgoURL. Cannot connect to database.")
	}

	logrus.Info("Connecting to mongo database at " + mgoURL)
	// TODO: Support private TLS certs
	MongoClient, err = mgo.Dial(mgoURL)
	if err != nil {
		logrus.Fatal("error connecting to mongo: " + err.Error())
	}
	logrus.Info("Connected to mongo.")
}
