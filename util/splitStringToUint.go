package util

import (
	"fmt"
	"strconv"
	"strings"
)

// splitStringToUint takes a string in the format "0-4" and returns the two parts as uint
func splitStringToUint(input string) (uint, uint, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input format")
	}

	// Convert the first part to uint
	first, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse first part: %v", err)
	}

	// Convert the second part to uint
	second, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse second part: %v", err)
	}

	return uint(first), uint(second), nil
}
