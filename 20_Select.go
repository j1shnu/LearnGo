// Select statement is used to choose multiple send/recv channel operations.
// Select statement blocks until one of the send/recv operation is ready.
// If multiple operations are ready, one of em is chosen at random.
// https://golangbot.com/select/

package main

import (
	"fmt"
	"time"
)

func downloader(t time.Duration, c chan bool) {
	time.Sleep(t * time.Second)
	c <- true
}

func main() {
	chan1 := make(chan bool)
	chan2 := make(chan bool)

	go downloader(8, chan1)
	go downloader(7, chan2)
        
        // This will exit without showing which one is completed 1st.
        fmt.Println("Methode 1:- ")
	select {
	case <-chan1:
		fmt.Println("Chan1 download Completed.")
	case <-chan2:
		fmt.Println("Chan2 download Completed.")
	default:
		fmt.Println("Download not Completed.")
	}

        // This will wait for one to complete.
        fmt.Println("Methode 2:- ")
        for {
          select {
          case <-chan1:
            fmt.Println("Chan1 download Completed.")
            return
          case <- chan2:
            fmt.Println("Chan2 download Completed.")
            return
          default:
            fmt.Println("Download In-Progres..")
            time.Sleep(3 * time.Second)
          }
        }

}
