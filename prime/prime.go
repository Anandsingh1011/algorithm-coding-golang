package main

import "fmt"

func main() {
	fmt.Println("***** Prime Number *****")

	fmt.Println("------- Check Prime Number ------")
	isPrimePrint(2)
	isPrimePrint(6)
	isPrimePrint(7)
	isPrimePrint(8)
	isPrimePrint(11)
	isPrimePrint(12)

	var num = 10
	fmt.Println("\n------- Print Fist ", num, "Prime Number -------")

	var i = 1

	for i < num {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
		i++
	}

	fmt.Println("")
}

func isPrime(num int) bool {

	var status bool = false

	if num == 1 || num == 2 {
		status = true
	} else if num%2 == 0 {
		status = false
	} else if num+1%2 == 0 {
		status = false
	} else {
		status = true
	}

	return status
}

func isPrimePrint(num int) bool {

	var status bool = isPrime(num)

	if status {
		fmt.Println(num, "is Prime Number ")
	} else {
		fmt.Println(num, "is NOT Prime Number ")
	}
	return status
}
