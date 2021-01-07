package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	done := make(chan bool)
	go func() {
		for {
			i++
			fmt.Println("i", i, <-ticker.C)

			if i == 5 {

				fmt.Println("stop", i)
				ticker.Stop()
				done <- true
			}
		}
	}()

	<-done

}
