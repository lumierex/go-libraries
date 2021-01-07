package groutine

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	v := strconv.Itoa(0)
	fmt.Println("v: ", v)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)

	go func() {
		intChan <- 1
	}()

	go func() {
		stringChan <- "hello world"
	}()

	select {
	case value := <-intChan:
		{
			fmt.Println(value)
		}
	case value := <-stringChan:
		{

			fmt.Println(value)
		}
	}
	fmt.Println("main 结束")
}



// isChannelFull receive slowly than write
func isChannelFull() {
	resultChan := make(chan string, 10)

	go selectWrite(resultChan)

	for value := range resultChan {

		fmt.Println("receive", value)
		time.Sleep(1000 * time.Millisecond)
	}

}

//
func selectWrite(result chan string) {

	for {
		select {
		case result <- "hello":
			{
				fmt.Println("write hello")
			}
		default:
			{
				log.Println("channel full")
			}

		}
		time.Sleep(500 * time.Millisecond)
	}

}
