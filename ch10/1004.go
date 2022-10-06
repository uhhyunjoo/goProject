package main

import "fmt"

func main() {
	day := "sunday"
	switch day {
	case "monday", "tuesday":
		fmt.Println("월화 : 수업가셈")
	case "wednesday", "thursday", "friday":
		fmt.Println("수목금 : 실습가셈")
	default:
		fmt.Println("주말이당")
	}
}
