package main

import "fmt"

func main() {
	// Declaring an array with size of 3
	var array1 [3]string = [3]string{"val1", "vat2", "val 3"}
	var array2 [4]int = [4]int{1, 2, 3, 4}
	var array3 = [3]int{1, 2, 3}
	array4 := [3]string{"a", "b", "c"}
	var array5 [3]int                              // Declaring without values
	array6 := [...]string{"me", "you", "us", "he"} // Declaring without size
	fmt.Println(array1, len(array1))               // Printing the array and it's size
	fmt.Println(array2, array3, array4, array5)
	array1[2] = "val4"
	fmt.Println(array1)
	fmt.Printf("%T, %v\n", array6, array6)

	// Declaring a slice. Size not required
	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("%T %v %d\n", slice1, slice1, len(slice1))

	var slice2 = append(slice1, 6)    // Appending to slice1
	var slice3 = append(slice2, 7, 8) // Appnding values to slice2
	fmt.Println(slice2, len(slice2))
	fmt.Println(slice3, len(slice3))

	// Slicing arrays and slices
	slice4 := array6[1:3]
	slice5 := array6[:3]
	slice6 := array6[3:]
	slice7 := array6[:] // Complete Array
	fmt.Println("Slicing array :-", slice4, slice5, slice6, slice7)

}
