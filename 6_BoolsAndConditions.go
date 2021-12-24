package main

import "fmt"

func main() {
	var age1 int8 = 25
	var age2 int8 = 30

	// Normal bool exprns
	fmt.Println(age1 == age2)
	fmt.Println(age1 < age2)
	fmt.Println(age1 >= age2)
	fmt.Println(age1 != age2)

	// If , else and nested if
	if age1 > age2 {
		fmt.Println(age1, "is greater than", age2)
	} else if age1 < age2 {
		fmt.Println(age1, "is lesser than", age2)
	} else {
		fmt.Println(age1, "is equal to", age2)
	}

	// Continue statement
	var array1 = []string{"jishnu", "jerin", "amma", "achan"}
	for index, value := range array1 {
		if index == 2 {
			fmt.Println("Skipping index 2")
			continue
		}
		fmt.Printf("Index of %v is %v\n", value, index)
	}

}
