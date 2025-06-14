package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  \nworld	 ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HelLO WoRLd",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("error: '%v' did not match expected '%v'", word, expectedWord)
			}
		}
	}
}
