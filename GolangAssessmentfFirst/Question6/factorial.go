package main

import "fmt"

func MyFact(num int) uint64 {

	if num >= 1 {

		return uint64(num) * MyFact(num-1)

	} else {

		return 1

	}

}

func main() {

	var num int

	fmt.Print("Enter Number:")

	fmt.Scanln(&num)

	fmt.Printf("Factorial of %d : %d\n", num, MyFact(num))

}
