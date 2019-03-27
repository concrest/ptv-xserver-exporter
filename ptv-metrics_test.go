package main

import (
	"fmt"
	"testing"
)

var timeConversions = []struct {
	fromPtv  float64
	expected float64
}{
	{27486515625000.0, 27486},
	{1236773546875000.0, 1236773.546875000},
}

func TestCpuTimeToSeconds(t *testing.T) {
	for _, testCase := range timeConversions {
		t.Run(fmt.Sprintf("%f", testCase.fromPtv), func(t *testing.T) {
			actual := cpuTimeToSeconds(testCase.fromPtv)
			difference := actual - testCase.expected

			if (difference > 1) || (difference < 0) {
				t.Errorf("Incorrect delta. Expect: %v. Actual: %v. Diff: %v", testCase.expected, actual, difference)
			}
		})
	}
}
