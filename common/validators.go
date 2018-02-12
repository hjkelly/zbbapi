package common

import "strings"

func StringIsEmpty(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}
