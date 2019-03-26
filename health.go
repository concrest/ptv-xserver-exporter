package main

import (
	"net/http"
	"time"
)

// GetProxyHandler creates a HTTP handler for the /proxy endpoint
func getHealthHandler() http.Handler {
	return &healthHandler{}
}

type healthHandler struct {
}

func (p *healthHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	defer timeTrack(time.Now(), "Health Check")

	// TODO. Define here what bad health looks like
	response.Header().Set("Content-Type", "text/plain")
	response.WriteHeader(200)
	response.Write([]byte("OK"))
}
