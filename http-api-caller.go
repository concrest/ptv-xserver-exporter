package main

import "errors"

// HTTPAPICaller is an APICaller for HTTP APIs
type HTTPAPICaller struct {
}

// GetBytes gets the bytes from a HTTP API endpoint
func (h *HTTPAPICaller) GetBytes(api string) ([]byte, error) {
	return nil, errors.New("Not implemented")
}
