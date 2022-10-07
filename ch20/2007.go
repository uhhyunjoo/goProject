package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf(">>>> 학생 나이 Stirng() 함수로 출력:%d", s.Age)
}

type Stringer interface {
	String() string
}

func PrintAge(stringer Stringer) {

	fmt.Println("type of stringer:", reflect.TypeOf(stringer)) // *main.Student

	// fmt.Println(">>>> 학생 나이 PrintAge() 함수로 출력: ", stringer.Age)

	s := stringer.(*Student)                     // 인터페이스 변수인 stringer 를 *Student 타입으로 타입 변환
	fmt.Println("type of s:", reflect.TypeOf(s)) // *main.Student
	fmt.Println(">>>> 학생 나이 PrintAge() 함수로 출력: ", s.Age)

}

func main() {
	s := &Student{15}

	fmt.Println(s.String()) // 아 이미 s 자체가 *Student 니까, receiver 로 s 써서 String() 함수 사용 가능

	PrintAge(s)
}
