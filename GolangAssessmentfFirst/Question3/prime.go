package main

import "fmt"

func main() {
	var num int = 0
	var num1 int = 0

	fmt.Print("enter the number:")
	fmt.Scanf("%d", num)

	for c := 2; c < num/2; c++ {
		if num%c == 0 {
			num1 = 1
			break
		}
	}
	if num1 == 1 {
		fmt.Print("given no is not prime")
	} else {
		fmt.Println("given no is prime")
	}
}
