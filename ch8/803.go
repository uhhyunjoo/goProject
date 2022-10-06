package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func printAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Cow {
		fmt.Println("꼬기오")
	} else {
		fmt.Println("...")
	}

}

func main() {
	printAnimal(Pig)
	printAnimal(Cow)
	printAnimal(Chicken)
}
