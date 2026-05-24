package main

import (
	"testing"
)

func TestGetDelimiterRuns(t *testing.T) {

	cases := []struct {
		input    string
		expected []DelimiterRun
	}{
		{input: "*****hello", expected: []DelimiterRun{{0, 3, true, false}}},
		{input: "**hello**", expected: []DelimiterRun{{0, 2, true, false}, {7, 2, false, true}}},
	}

	for _, testCase := range cases {

		t.Run(testCase.input, func(t *testing.T) {
			results := GetDelimiterRuns(testCase.input)

			for i, r := range results {
				if r != testCase.expected[i] {
					t.Errorf("got %v, want %v", results, testCase.expected)
				}
			}
		})
	}

}
