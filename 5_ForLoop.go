package main

import "fmt"

func main() {
	// Type :- 1
	i := 0
	for i <= 4 {
		fmt.Println("Value of i is :-", i)
		i++
	}

	// Type :- 2
	for i := 5; i <= 10; i++ {
		fmt.Println("Value of i is :-", i)
	}

	// Type :- 3 - Using range
	myArray := [5]string{"tom", "micky", "jerry", "bean", "donald"}
	for index, value := range myArray {
		fmt.Printf("Index of %v is %v\n", value, index)
	}

}
