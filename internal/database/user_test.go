package database

import "fmt"

func RunTest() {
	me := newUser().
		Id(122).
		Username("MASTER_TEST").
		Fullname("Master Test").
		Phone("999-999-999").
		Email("example@test.com").
		Account(12909292).
		Build()

	fmt.Println(me)
}
