package main

import "fmt"

func main() {
	var num1, num2, sum int

	fmt.Println("Enter a number:")
	fmt.Scan(&num1)
	fmt.Println("Enter a number 2 :")
	fmt.Scan(&num2)

	sum = num1 + num2
	fmt.Println("The sum is : ", sum)

	sum = num1 - num2
	fmt.Println("The difference is : ", sum)

	sum = num1 * num2
	fmt.Println("The product is : ", sum)

	sum = num1 / num2
	fmt.Println("The quotient is : ", sum)

	sum = num1 % num2
	fmt.Println("The remainder is : ", sum)
}
