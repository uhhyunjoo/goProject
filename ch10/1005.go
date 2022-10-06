package main

import "fmt"

func main() {
	// temp := 9
	// switch true { // 비굣값 쓰지 않고, if문 처럼 조건문 검사하는 경우 true 생략해도 됨
	// case temp < 10 || temp > 30:
	// 	fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	// case temp >= 10 && temp < 20:
	// 	fmt.Println("약간 추움 겉옷 ㄱ")
	// case temp >= 15 && temp < 25:
	// 	fmt.Println("야외활동 굿")
	// default:
	// 	fmt.Println("따뜻~")
	// }

	temp := 5
	switch true { // 비굣값 쓰지 않고, if문 처럼 조건문 검사하는 경우 true 생략해도 됨
	case temp < 10 || temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
		fallthrough
	case temp < 20:
		fmt.Println("약간 추움 겉옷 ㄱ")
	case temp <= 15:
		fmt.Println("야외활동 굿")
	default:
		fmt.Println("따뜻~")
	}
}
