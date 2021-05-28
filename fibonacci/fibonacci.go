package main

import (
	"fmt"
)

func main() {

	var num = 10
	fmt.Println("Print first ", num, " Fibonacci number :")
	fbonacci(num)
	fmt.Println("")
}

func fbonacci(num int) {

	if num == 1 {
		fmt.Print(num)
	} else if num == 2 {
		fmt.Print(1, 1)
	} else {

		fmt.Print(1, 1, " ")

		var count = 2
		var n1, n2, n3 int = 1, 1, 0

		for count < num {

			n3 = n1 + n2
			fmt.Print(n3, " ")
			n1 = n2
			n2 = n3

			count++
		}
	}

}
