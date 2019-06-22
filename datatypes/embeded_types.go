package main

import (
	"fmt"
)


type user struct {
	name string
	email string
}

type notifier interface {
	notify()
}


func (u *user) notify() {
	fmt.Println("Sending email to user:", u.name)
}

type admin struct {
	user
	level string
}

func main(){
	fmt.Println("This is a test for embeded type system in golang")


	ad := admin{
		user: user{
			name:"Ronglu Cao",
			email: "ronglu.cao@oracle.com",
		},
		level: "3",
	}

	ad.notify()
	ad.user.notify()

	fmt.Println("Name from ad directly:", ad.name)
}