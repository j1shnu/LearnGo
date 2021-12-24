package main

import "fmt"

func updateString(s string) {
	s = "Upadated value"
}

func updateStringNew(s *string) {
	*s = "Updated value"
}

func main() {
	print := fmt.Println
	// while passing in a func :- grp1 will make a copy and pass and grp2 will pass the copy with same mem location.
	// var group1 = []string{"strings", "ints", "floats", "booleans", "arrays", "structs"}
	// var group2 = []string{"slices", "maps", "functions"}

	string1 := "nice string"
	print(string1)
	updateString(string1)
	print(string1)
	updateStringNew(&string1)
	print(string1)

	// Declaring a pointer with string2's mem addr
	string2 := "good string"
	// Two methods to declare pointer
	pointer1 := &string2
	var pointer2 *string = &string2

	print(*pointer1)
	updateStringNew(pointer1)
	print(*pointer1)
	print(*pointer2)

}
