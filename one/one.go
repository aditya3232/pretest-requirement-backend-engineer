package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	input := "121"
	fmt.Println(isPalindrome(input))

	input = "semangat"
	fmt.Println(isPalindrome(input))
}
