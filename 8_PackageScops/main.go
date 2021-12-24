package main

import "fmt"

func main() {
	fmt.Println(sayHi("jishnu"))
	for _, n := range names {
		fmt.Println(sayHi(n))
	}
}
