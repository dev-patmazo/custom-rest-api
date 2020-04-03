package main

import (
	"fmt"
	"net/http"

	"rest-api/config"

	"goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	mux := goji.NewMux()
	fmt.Println("Initializing configuration..")
	cfg := config.InitializeConfig()
	fmt.Printf("%+v", cfg.Server.Port)

	//fmt.Println("We are actually running at port :" + cfg.Server.Port)
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	http.ListenAndServe("localhost:8080", mux)
}
