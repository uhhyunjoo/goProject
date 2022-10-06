package main

import "fmt"

func main() {

	a := 3
	var b float64 = 3.5

	// var c int = b (float 인 b 를 int 인 c 에 대입하려고 하니까 에러)
	// d := a * b (int 인 a 랑 float 인 b 를 곱하려고 해서 에러... 다른 타입이라 그렇대;;;;)

	var e int64 = 7
	// f := a * e (int 인 a 랑 int64 인 e 를 연산하려고 해서 에러... 같은 int 인데도 다른 타입이라 그렇대;;;;;)

	// fmt.Println(a, b, c, d, e, f)

	fmt.Println(a, b, e)
}
