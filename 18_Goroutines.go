// Goroutines are funcs or methods that run concurrently with other funcs or methods.
package main

import (
  "fmt"
  "time"
)

func testGo(){
  fmt.Println("From Goroutines func")
}
func routineFunc(n int){
  for i:=0; i<=n; i++{
    fmt.Println(i)
  }
}
func main(){
  go testGo() // Need to wait to get result otherwise program will exit without the output.
  go routineFunc(10)
  fmt.Println("From main func")
  time.Sleep(1 * time.Millisecond) // Wait for goroutines or can use channels
}
