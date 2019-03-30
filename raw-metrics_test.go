package main

import (
	"io/ioutil"
	"testing"
)

func TestParseExampleXMapJSON(t *testing.T) {
	bytes, err := ioutil.ReadFile("testData/example-xmap.json")
	if err != nil {
		t.Errorf("File read error: %v", err)
		return
	}

	actual, err := NewRawMetrics(bytes)
	if err != nil {
		t.Errorf("NewRawMetrics error: %v", err)
		return
	}

	if actual.ServiceName != "PTV xMap Server" {
		t.Errorf("Unexpected ServiceName '%s'", actual.ServiceName)
	}

	if len(actual.Instances) != 2 {
		t.Errorf("Expected 2 Instances, but was %v", len(actual.Instances))
	}

	t.Logf("Success: %+v", actual)
}
