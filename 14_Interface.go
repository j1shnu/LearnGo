// Interfaces are named collections of method signatures.
// https://medium.com/rungo/interfaces-in-go-ab1601159b3a
package main

import "fmt"

type girl struct {
	name  string
	class string
}

type boy struct {
	name  string
	class string
}

func (g girl) isIn() string {
	return fmt.Sprintf("The girl named %v is in %v class now.\n", g.name, g.class)
}

func (b boy) isIn() string {
	return fmt.Sprintf("The boy named %v is in %v class now.\n", b.name, b.class)
}

// Creating an interface
// Define :- Any type or struct that has isIn methode and returns string is of type whereAre.
type whereAre interface {
	isIn() string
}

func whereIs(w whereAre) string {
	return w.isIn()
}

func main() {
	student1 := boy{"Jishnu", "CS"}
	student2 := girl{"Sheela", "Chemistry"}

	// Normal Strut func / Receiver func call
	fmt.Print(student1.isIn())
	fmt.Println(student2.isIn())

	student3 := boy{"Arun", "Biology"}

	// Using Interface
	// Method: 1
	students := []whereAre{student1, student2, student3}
	for _, student := range students {
		fmt.Print(student.isIn())
	}
	fmt.Println()

	// Method: 2
	fmt.Print(whereIs(student1))
	fmt.Print(whereIs(student2))
	fmt.Print(whereIs(student3))
	fmt.Println()

	// Method: 3
	var w whereAre = boy{"James", "Physics"}
	fmt.Print(w.isIn())
}
