package main

import "fmt"

func main() {
	// Strings
	// Declaring Varibles with values
	var name1 string = "Jishnu"
	var name2 = "Prasad"
	name3 := "KP"    // Can't use outside the function
	var name4 string // Declaring without value

	fmt.Println(name1, name2, name3, name4)

	// Updating Values
	name1 = "Jerin"
	name4 = "GoLang"

	fmt.Println(name1, name4)

	// Numbers
	// https://go.dev/ref/spec#Numeric_types  -  Numeric types and it's range in GO
	var age1 int = 25
	var age3 int8 = 10
	age4 := 35
	var age5 float64

	fmt.Println(age1, age3, age4, age5)
}
