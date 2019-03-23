package main

import "io/ioutil"

// FileAPICaller is a testing APICaller for getting files
type FileAPICaller struct {
	Filename string
}

// GetBytes gets the bytes from the filesystem
func (h *FileAPICaller) GetBytes(api string) ([]byte, error) {
	return ioutil.ReadFile(h.Filename)
}
