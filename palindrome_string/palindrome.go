package main

import (
	"fmt"
)

func main() {
	fmt.Println("Palindrome String")

	isPalindrome("anand")
	isPalindrome("nitin")
	isPalindrome("ancnc")
}

func isPalindrome(str string) bool {
	var status = true

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			status = false
		}
	}

	if status {
		fmt.Println(str, " is a Palindrome string")
	} else {
		fmt.Println(str, " is NOT a Palindrome string")
	}
	return status
}
