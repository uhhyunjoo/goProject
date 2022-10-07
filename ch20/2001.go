package main

import "fmt"

/*

type 인터페이스명 interface{

	// {} 안에 메서드 집합

	Fly()
	Walk(distance int) int
}

*/

type Stringer interface {
	printInfo() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) printInfo() string {
	return fmt.Sprintf("안녕! 내 이름은 %s이고 %d살이야", s.Name, s.Age)
}

func main() {
	student := Student{"곰곰", 100}

	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.printInfo())

	stringer2 := Student{"떙떙", 10}

	fmt.Printf("%s\n", stringer2.printInfo())

}
