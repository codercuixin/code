package main 

import "fmt"

type notifier interface{
	notify()
}

type user struct{
	name string
	email string 
}

func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}



type admin struct{
	user
	level string 
}




func sendNotification(n notifier){
	n.notify()
}

func main(){
	// u := user{
	// 	name: "user",
	// 	email: "user@email.com",
	// }
	a := admin{
		user: user{
		 name: "amdin",
		 email: "amdin@email.com",
		},
		level: "amdin@email.com",
	}

	a.user.notify()
	a.notify()
}

func (a *admin) notify(){
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}