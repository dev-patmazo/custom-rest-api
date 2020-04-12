package main

import (
	"fmt"
	"net/http"
	"rest-api/config"

	"github.com/sirupsen/logrus"
	"goji.io"
	"goji.io/pat"
)

var cfg config.ConfigTemplate

func init() {
	config.SetLogConfig()
	config.SetEnvConfig()
	config.SetDBCOnfig()
	cfg = config.GetEnvInfo()
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	logrus.Info("Application is up and running on port : " + cfg.Server.Port)
	http.ListenAndServe("localhost:8080", mux)
}
