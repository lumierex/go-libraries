package groutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var l sync.Mutex


func normalAdd() {
	x++
	wg.Done()
}

func mutexAdd() {

	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func atomicConcurrent() {
	start := time.Now()
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

func normalConcurrent() {
	start := time.Now()
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go normalAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
