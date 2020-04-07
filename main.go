package main

import (
	"fmt"
	"net/http"
	"rest-api/config"

	"goji.io"
	"goji.io/pat"
)

var logger = config.Logger()

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	mux := goji.NewMux()
	logger.Info("Initializing configuration..")
	cfg := config.InitializeConfig()
	logger.Info("Application works well and running on port :" + cfg.Server.Port)

	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	http.ListenAndServe("localhost:8080", mux)
}
