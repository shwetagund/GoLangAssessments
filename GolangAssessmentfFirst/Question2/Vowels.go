package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the char:")
	ch, _ := reader.ReadByte()

	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		fmt.Printf("%c is a vowels char\n", ch)
	} else {
		fmt.Printf("%c is not vowels\n", ch)
	}

}
