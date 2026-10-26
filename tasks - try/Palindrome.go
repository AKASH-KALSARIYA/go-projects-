package main

import "fmt"

func main() {
	var num, temp, reverse int

	fmt.Print("Enter Number: ")
	fmt.Scan(&num)

	temp = num

	for num > 0 {
		digit := num % 10
		reverse = reverse*10 + digit
		num = num / 10
	}

	if temp == reverse {
		fmt.Println("Palindrome Number")
	} else {
		fmt.Println("Not Palindrome")
	}
}
