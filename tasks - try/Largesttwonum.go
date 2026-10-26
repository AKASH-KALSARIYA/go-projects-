package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("enter a num1: ")
	fmt.Scan(&num1)
	fmt.Print("enter a num2: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Print("num 1 is bigger!!!")
	} else if num1 < num2 {
		fmt.Println("num 2 is biggest!!!")
	} else {
		fmt.Println("num is qual!!!")
	}
}
