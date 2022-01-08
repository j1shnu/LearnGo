// Channels are like pipes using which Goroutines communicate.
// Each channel has a type associated with it.

package main

import "fmt"

func testFunc1(n int, c chan bool){
  for i:=0; i<=n; i++ {
    fmt.Printf("%v, ", i)
  }
  c <- true
}

func getData(c chan string){
  c <- "Chunk of Data"
}

func main(){
  // Creating a channel
  var chan1 chan int //Method 1
  done := make(chan bool) //Method 2
  fmt.Println(chan1) 
  dataChan := make(chan string)

  go testFunc1(20, done)
  go getData(dataChan)
  
  <-done 
  fmt.Printf("\n\n%v\n", <-dataChan) // Will print the string passed to the channel
}
