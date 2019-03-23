package main

import (
	"io/ioutil"
	"testing"
)

func TestScrape(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("testData/example-xlocate.json")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
		return
	}

	fakeCaller := &fakeAPICaller{
		Bytes: fileBytes,
		Error: nil,
	}

	scraper := NewScraper("/foo/bar", fakeCaller)
	rawMetrics, err := scraper.Scrape()

	if err != nil {
		t.Errorf("Error in Scrape: %v", err)
	}

	if rawMetrics.ServiceName != "PTV xLocate Server" {
		t.Errorf("Unexpected ServiceName: %s", rawMetrics.ServiceName)
	}
}

type fakeAPICaller struct {
	Error error
	Bytes []byte
}

func (c *fakeAPICaller) GetBytes(api string) ([]byte, error) {
	return c.Bytes, c.Error
}
