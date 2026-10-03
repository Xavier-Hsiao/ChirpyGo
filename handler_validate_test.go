package main

import "testing"

func TestValidateProfane(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		badWords []string
		expected string
	}{
		{
			name:     "no profane words",
			input:    "I had something interesting for breakfast",
			badWords: []string{"kerfuffle", "sharbert", "fornax"},
			expected: "I had something interesting for breakfast",
		},
		{
			name:     "single profane word",
			input:    "I need a kerfuffle to sleep",
			badWords: []string{"kerfuffle", "sharbert", "fornax"},
			expected: "I need a **** to sleep",
		},
		{
			name:     "case insensitive match",
			input:    "Fornax is a word",
			badWords: []string{"kerfuffle", "sharbert", "fornax"},
			expected: "**** is a word",
		},
		{
			name:     "word with punctuation is not replaced",
			input:    "Sharbert! is different",
			badWords: []string{"kerfuffle", "sharbert", "fornax"},
			expected: "Sharbert! is different",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := validateProfane(tc.input, tc.badWords)
			if actual != tc.expected {
				t.Errorf("validateProfane(%q, %v) = %q; expected %q", tc.input, tc.badWords, actual, tc.expected)
			}
		})
	}
}
