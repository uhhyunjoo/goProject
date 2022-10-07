package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func (a *account) withdrawToStringMethod(amount int) string {
	a.balance -= amount

	return string(rune(a.balance)) + "입니당"
}

func main() {
	a := &account{100}

	withdrawFunc(a, 30)

	a.withdrawMethod(30)

	s := a.withdrawToStringMethod(20)

	fmt.Printf("%d \n", a.balance)

	fmt.Println(s)

}
