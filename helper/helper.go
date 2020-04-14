package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	//respWriter http.ResponseWriter
	response interface{}
	message  = make(map[string]interface{})
)

func RequestBody(req *http.Request, respWriter http.ResponseWriter) map[string]interface{} {

	if body, err := ioutil.ReadAll(req.Body); err != nil {
		message["code"] = http.StatusBadRequest
		message["status"] = http.StatusText(http.StatusBadRequest)
		Response(message, respWriter)
	} else {
		if err := json.Unmarshal(body, &message); err != nil {
			message["code"] = http.StatusBadRequest
			message["status"] = http.StatusText(http.StatusBadRequest)
			Response(message, respWriter)
		} else {
			if len(message) != 0 {
				return message
			}
			message["code"] = http.StatusBadRequest
			message["status"] = http.StatusText(http.StatusBadRequest)
			Response(message, respWriter)
		}
	}

	return message

}

func Response(data interface{}, respWriter http.ResponseWriter) {

	if response, err := json.Marshal(data); err != nil {
		log.Debug(err.Error())
		return
	} else {
		respWriter.Write(response)
		return
	}

	return

}
