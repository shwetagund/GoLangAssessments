package main

import "fmt"

func main() {
	var myString1 = "shweta\tgund"
	fmt.Println("Using tab:", myString1)

	var myString2 = "shweta\ngund"
	fmt.Println("using new line string:", myString2)

	var myString3 = "Mr.ABC "
	fmt.Printf("%q", myString3)

}
