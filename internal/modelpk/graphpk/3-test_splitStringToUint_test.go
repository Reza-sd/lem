package graphpk

import (
	//"reflect"
	"testing"
)

func TestSplitStringToParts(t *testing.T) {
	t.Skip()
	tests := []struct {
		name          string
		input         string
		expectedPart1 string
		expectedPart2 string
		expectedError bool
	}{
		{
			name:          "Valid input",
			input:         "part1-part2",
			expectedPart1: "part1",
			expectedPart2: "part2",
			expectedError: false,
		},
		{
			name:          "Missing hyphen",
			input:         "invalidinput",
			expectedPart1: "",
			expectedPart2: "",
			expectedError: true,
		},
		{
			name:          "Too many hyphens",
			input:         "part1-part2-part3",
			expectedPart1: "",
			expectedPart2: "",
			expectedError: true,
		},
		{
			name:          "Empty input",
			input:         "",
			expectedPart1: "",
			expectedPart2: "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1, part2, err := splitStringToParts(tt.input)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if part1 != tt.expectedPart1 {
				t.Errorf("Expected part1 to be %q, but got %q", tt.expectedPart1, part1)
			}

			if part2 != tt.expectedPart2 {
				t.Errorf("Expected part2 to be %q, but got %q", tt.expectedPart2, part2)
			}
		})
	}
}
