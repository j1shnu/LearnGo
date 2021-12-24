package main

// https://pkg.go.dev/fmt
import "fmt"

func main() {
	// print without new line. Need to specify newline(\n)
	fmt.Printf("Hello ")
	fmt.Printf("World..!\n")
	fmt.Printf("Hi..")
	fmt.Printf("Howdy..?")

	// print with newline
	fmt.Println("Hello..")
	fmt.Println("World..!")

	var name string = "Jishnu"
	var age int8 = 25

	// Printing type of variable
	fmt.Printf("Type of name and age variables are %T and %T\n", name, age)
	// Printing variable values
	fmt.Println("My name is", name, "and I am", age)
	fmt.Printf("My name is %v and I am %v\n", name, age) // %v for variable
	fmt.Printf("My name is %s and I am %d\n", name, age) // Specifying variables type for printing

	var sprintString = fmt.Sprintf("Sprintf :- My name is %v and I am %v\n", name, age)
	fmt.Println("Saved string =>", sprintString)
}
