package cstring

import (
	"strings"
)

func CountPrefix(s, substr string) int {
	if s == "" || substr == "" {
		return 0
	}
	n := 0
	for strings.HasPrefix(s, substr) {
		s = s[len(substr):]
		n++
	}
	return n
}
