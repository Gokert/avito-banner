package httpResponse

import (
	"avito-banner/pkg/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func SendResponse(w http.ResponseWriter, r *http.Request, response *models.Response, firstTime time.Time, log *logrus.Logger) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Error("Send response error: ", err)
		response.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	timeNow := time.Now()
	responseTime := timeNow.Sub(firstTime)

	log.Infof("Time: %v, Response time: %v, Address: %s, Method: %s, Status: %d, URL: %s", timeNow, responseTime, r.RemoteAddr, r.Method, response.Status, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error("Failed to send response: ", err.Error())
	}
}
