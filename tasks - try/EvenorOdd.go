package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter  a num:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("num is even!!")
	} else {
		fmt.Print("num is odd!!!")
	}
}
