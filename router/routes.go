package router

import (
	"rest-api/controller"

	"goji.io"
	"goji.io/pat"
)

func EndPoints() (mux *goji.Mux) {
	mux = goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), controller.Hello)

	return mux
}
