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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Lanturn is Cool!",
			expected: []string{"lanturn", "is", "cool!"},
		},
		{
			input:  "charizard is LAME",
			expected: []string{"charizard", "is", "lame"},
		},

	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Slice length incorrect")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word incorrect")
				continue
			}
		}
	}
}