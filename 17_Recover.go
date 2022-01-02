// recover is used to regain control of a panicking program
// Usefull only when called under defer
// https://golangbot.com/panic-and-recover/
package main

import "fmt"

func main(){
  fmt.Println("Start")
  slice1 := []int{1, 2, 3}
  // recover panick using anonymouse func
  defer func(){
    if err := recover(); err != nil {
      fmt.Println("Recovered error:- ", err)
    }
  }()
  // This'll panic (index out of range)
  arrayVal1 := slice1[4]
  fmt.Println("val1", arrayVal1)
  fmt.Println("Done")
}
