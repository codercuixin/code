package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct{
	person user 
	level string 
}

type Duration int64 

func main() {
	// struct1()
	// structAdmin()
	typeDuration()
}



func typeDuration(){
	var dur Duration
	dur = Duration(1000)
	fmt.Println(dur)
}


func structAdmin(){
	fred := admin{
		person: user{
			name : "Lisa",
			email: "lisa@email.com",
			ext :123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)
}

func struct1(){
	var lisa user
	lisa = user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(lisa)

	bill := user{"Bill", "bill@email.com", 123, true}
	fmt.Println(bill)
}