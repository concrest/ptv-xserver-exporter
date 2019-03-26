package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// HTTPAPICaller is an APICaller for HTTP APIs
type HTTPAPICaller struct {
}

// GetBytes gets the bytes from a HTTP API endpoint
func (h *HTTPAPICaller) GetBytes(api string) ([]byte, error) {
	defer timeTrack(time.Now(), "HTTPAPICaller.GetBytes")

	log.WithFields(log.Fields{
		"api": api,
	}).Debug("Calling HTTP API")

	resp, err := http.Get(api)
	if err != nil {
		log.WithFields(log.Fields{
			"api": api,
			"err": err,
		}).Warn("http.Get returned an error")

		return []byte{}, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.WithFields(log.Fields{
			"api":        api,
			"err":        err,
			"statusCode": resp.StatusCode,
		}).Warn("Error reading response body")

		return []byte{}, err
	}

	responseBodyText := fmt.Sprintf("%s", responseBody)

	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"api":        api,
			"statusCode": resp.StatusCode,
			"body":       responseBodyText,
		}).Warn("Server returned non-200 response")

		return []byte{}, fmt.Errorf("GET %s returned HTTP %d", api, resp.StatusCode)
	}

	log.WithFields(log.Fields{
		"api":  api,
		"body": responseBodyText,
	}).Debug("Received response data")

	return responseBody, nil
}
