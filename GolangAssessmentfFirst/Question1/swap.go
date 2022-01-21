package main

import "fmt"

func main() {

	num1 := 10
	num2 := 20
	fmt.Printf("before swap num1=%d,num2=%d\n", num1, num2)

	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	fmt.Printf("after swapping num1=%d,num2=%d\n", num1, num2)

}
