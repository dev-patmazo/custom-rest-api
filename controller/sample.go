package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"goji.io/pat"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	log.WithField("Hello", name).Info()
}
