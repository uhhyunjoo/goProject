package main

import "fmt"

func main() {
	// var 변수명 [요소개수]타입
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
