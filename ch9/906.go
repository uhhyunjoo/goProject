package main

import "fmt"

func getMyAge() (int, bool) {
	return 10, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("Your are young", age)
		fmt.Println("ok:", ok, "age:", age < 20, "ok && age:", ok && age < 20)
	} else if ok && age < 30 {
		fmt.Println("Niceage", age)
	} else if ok {
		fmt.Println("Niceage", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is", age) : undeclared age
}
