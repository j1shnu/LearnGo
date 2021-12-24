package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Methode 1
	fmt.Print("Enter Your First Name:- ")
	var fname string
	fmt.Scanf("%s", &fname)
	fmt.Println("'", fname, "'")

	// Methode 2
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Second Name:- ")
	sname, _ := reader.ReadString('\n')
	sname = strings.TrimSpace(sname)
	fmt.Println("'", sname, "'")
}
