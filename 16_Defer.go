package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func main(){
  resp, err := http.Get("http://ifconfig.me")
  if err != nil {
    log.Fatal(err)
  }
  // Defer will execute just before exiting the main func
  defer resp.Body.Close()
  out, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", out)
}
