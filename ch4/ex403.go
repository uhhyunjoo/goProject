package main

import "fmt"

func main() {
	var a int = 3
	var b int // 0 이 자동으로 들어감
	var c = 4 // 정수 4로 저장됨
	d := 5    // 선언 대입문 := 을 사용해 var 키워드와 타입 생략 // 같은 맥락으로, 정수 5로 들어감
	// e = 6 // 이건 에러 난다.

	fmt.Println(a, b, c, d)
}
