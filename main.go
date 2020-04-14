package main

import (
	"net/http"
	"rest-api/config"
	"rest-api/middleware"
	"rest-api/router"

	log "github.com/sirupsen/logrus"
	"goji.io"
)

var (
	cfg config.ConfigTemplate
	mux *goji.Mux
)

func init() {
	config.SetEnvConfig()
	config.SetDBCOnfig()
	config.SetLogConfig()
	mux = router.EndPoints()
	cfg = config.GetEnvInfo()

	//test
	middleware.Tokenizer()
	middleware.Detokenizer()
}

func main() {
	log.Info("Application is up and running on port : " + cfg.Server.Port)
	http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, mux)
}
