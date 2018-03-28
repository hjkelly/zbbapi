package common

import "strings"

// StringIsEmpty returns true if the string contains any non-whitespace characters.
func StringIsEmpty(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}
