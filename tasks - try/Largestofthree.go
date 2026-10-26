package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println("Largest =", a)
	} else if b >= c {
		fmt.Println("Largest =", b)
	} else {
		fmt.Println("Largest =", c)
	}
}
