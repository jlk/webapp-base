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

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)

	})
}
