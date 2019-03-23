package main

import (
	"fmt"
	"net/http"
)

// GetProxyHandler creates a HTTP handler for the /proxy endpoint
func GetProxyHandler(api string) http.Handler {
	return &proxyHandler{
		scraper: NewScraper(api, &HTTPAPICaller{}),
		//scraper: NewScraper(api, &FileAPICaller{
		//	Filename: "testData/example-xmapmatch.json",
		//}),
	}
}

type proxyHandler struct {
	scraper *Scraper
}

func (p *proxyHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	metrics, err := p.scraper.Scrape()
	if err != nil {
		response.Header().Set("Content-Type", "text/plain")
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	response.Header().Set("Content-Type", "text/plain")
	response.WriteHeader(200)
	response.Write([]byte(fmt.Sprintf("%+v", metrics)))
}
