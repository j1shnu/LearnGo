package main

import "fmt"

// cant access publicly if declared in a func
var names = []string{"kuttu", "appu", "monna", "bindu", "prasad", "radha"}

func sayHi(name string) string {
	return fmt.Sprintf("Hi %v ..!", name)
}
