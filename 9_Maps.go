package main

import "fmt"

// https://www.geeksforgeeks.org/golang-maps/
// Map is like key:value pairs. Like dict in python
func main() {
	// Declaring a map with values
	var map1 map[string]int = map[string]int{
		"jishnu": 25,
		"jerry":  21,
		"micky":  23,
	}
	fmt.Println(map1)

	// Declaring empty map
	map2 := make(map[int]string)
	var map3 map[string]int
	fmt.Println(map2, map3)

	// Updating map
	map1["jishnu"] = 26

	// Adding key value
	map1["donald"] = 48

	fmt.Println(map1)

	// Looping through maps
	for key, value := range map1 {
		fmt.Println(key, ":-", value)
	}

	// Deleting from map
	delete(map1, "jishnu")
	fmt.Println(map1)
}
