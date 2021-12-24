// Struct is user-defained type. egs of types are string, int, func...etc
// https://www.golangprograms.com/go-language/struct.html

package main

import "fmt"

// Declaring a struct
type user struct {
	fullname string
	email    string
	age      int8
}

func createUser(name string, email string, age int8) user {
	//
	newUser := user{
		fullname: name,
		email:    email,
		age:      age,
	}
	// anotherWay := user{name, email, age}  - Simple way to create instance of strauct

	// Other way to create instance of strauct
	/* var otherWay user
	otherWay.fullname = name
	otherWay.email = email
	otherWay.age = age */
	return newUser
}
func main() {
	jishnu := createUser("jishnu", "jishnu@mail.com", 25)
	fmt.Println(jishnu)
}
