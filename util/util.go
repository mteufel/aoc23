package util

import "unicode"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
