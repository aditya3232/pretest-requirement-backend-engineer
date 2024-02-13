package one

import (
	"strings"
)

func IsPalindrome(d string) bool {

	s := strings.ToLower(d)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
