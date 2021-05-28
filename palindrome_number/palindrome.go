package main

import (
	"fmt"
)

func main() {
	fmt.Println("Palindrome Number")

	isPalindrome(1234)
	isPalindrome(456654)
	isPalindrome(1001)
	isPalindrome(1002001)
}

func isPalindrome(num int) bool {
	var status = true

	var remainder, temp int
	var reverse int = 0

	temp = num

	for {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10

		if num == 0 {
			// Break the loop
			break
		}
	}

	if temp == reverse {
		fmt.Println(temp, " is a Palindrome number")
		status = true
	} else {
		fmt.Println(temp, " is NOT a Palindrome number")
		status = false
	}
	return status
}
