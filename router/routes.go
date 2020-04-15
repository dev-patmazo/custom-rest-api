package router

import (
	"rest-api/controller"
	"rest-api/middleware"

	"goji.io"
	"goji.io/pat"
)

func EndPoints() (mux *goji.Mux) {

	//public := goji.NewMux() For open apis

	private := goji.NewMux()
	private.Use(middleware.Interceptor)
	private.HandleFunc(pat.Get("/hello/:name"), controller.Hello)

	return mux
}
