package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//var  a = uint8(12)
	////sync.Pool{New: }
	//fmt.Println(a)

	var t time.Time
	//fmt.Println()
	fmt.Println(t.Add(time.Second * 2).Before(time.Now()))

	fmt.Println("time: Add", t.Add(time.Second*2))

	fmt.Println("sub: ", t.Sub(time.Now()))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)

	defer cancel()
	go helloHandle(ctx, 2000*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("Hello Handle request timeout", ctx.Err())
	}

}

func helloHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
