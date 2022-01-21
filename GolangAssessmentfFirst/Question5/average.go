package main

import "fmt"

func main() {
	array := [11]int{2, 4, 6, 8, 34, 5, 7, 3, 5, 67, 4}
	length := 11
	sum := 0
	for i := 0; i < length; i++ {
		sum += (array[i])
	}
	average := sum / length
	fmt.Println("sum=", sum, "average=", average)

}
