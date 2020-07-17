package main

import (
	"log"
	"net/http"

	"github.com/jlk/webapp-base/server/data"
	"github.com/sirupsen/logrus"
)

func main() {

	log.Println("Now server is running on port " + data.Config.GetString("ListenPort"))
	err := http.ListenAndServe(":"+data.Config.GetString("ListenPort"), nil)
	if err != nil {
		logrus.Errorf("Error when attempting to start network server: %s", err.Error())
	}
}
