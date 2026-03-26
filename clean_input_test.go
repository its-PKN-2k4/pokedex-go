package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "Tepig  \t SNIVY \n oshawott",
			expected: []string{"tepig", "snivy", "oshawott"},
		},
		{
			input:    "ZEKROM \n RESHIRAM",
			expected: []string{"zekrom", "reshiram"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if !reflect.DeepEqual(len(c.expected), len(actual)) {
			t.Errorf("Expected: %v; Got: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if !reflect.DeepEqual(expectedWord, word) {
				t.Errorf("Expected word: %v; Got: %v", expectedWord, word)
			}
		}
	}
}
