package domain

import (
	"fmt"
	"time"
)

type User struct {
	Id       uint64
	NickName string
	Time     time.Duration
}

func (u User) ShowName() {
	fmt.Printf("My name is: %s", u.NickName)
}
