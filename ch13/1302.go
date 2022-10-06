package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"김익명", "ikmyeong", 25}

	vip := VIPUser{User{"김손님", "ing", 100}, 2, 30000}

	fmt.Println(user.Name, user.ID, user.Age)
	fmt.Println(vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)
}
