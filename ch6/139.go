package main

import "fmt"

func main() {
	var x1 int8 = 2
	var x2 int8 = 10

	fmt.Printf("%d&^%d = %5d,  \t %08b\n", x2, x1, x2 & ^x1, uint(x2 & ^x1))

}
