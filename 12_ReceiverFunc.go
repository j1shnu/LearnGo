package main

import "fmt"

type user struct {
	fullname string
	email    string
	age      int
}

func createUser(name string, email string, age int) user {
	return user{name, email, age}
}

// Reciever func
func (s user) print() {
	fmt.Printf("%v's email is %v and is %v years old.\n", s.fullname, s.email, s.age)
}

func main() {
	jishnu := createUser("jishnu", "jishnu@email.com", 25)
	fmt.Println(jishnu)
	jishnu.print()
}
