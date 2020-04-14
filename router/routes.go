package router

import (
	"rest-api/controller"
	"rest-api/middleware"

	"goji.io"
	"goji.io/pat"
)

func EndPoints() (mux *goji.Mux) {
	mux = goji.NewMux()
	mux.Use(middleware.Interceptor)
	mux.HandleFunc(pat.Get("/hello/:name"), controller.Hello)

	return mux
}
