package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8
	B int8
	C int8
	D int
	E int8
}

func main() {
	user := User{1, 2, 3, 4, 5}

	v := "90"
	fmt.Println(unsafe.Sizeof(user))
	fmt.Println(unsafe.Sizeof(v))
}
