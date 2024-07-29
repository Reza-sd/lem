package graphpk

import (
	"fmt"
	//"strconv"
	"strings"
)

// splitStringToUint takes a string in the format "0-4" and returns the two parts as uint
func splitStringToUint(input string) (string, string, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid input format")
	}


	return parts[0], parts[1], nil
}
