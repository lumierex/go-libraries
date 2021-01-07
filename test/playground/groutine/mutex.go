package groutine

import (
	"fmt"
	"sync"
	"time"
)

var (
	a      int64
	x      int64
	y      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func rwLockRead() {
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func rwLockWrite() {
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	x = x + 1
	rwlock.Unlock()
	wg.Done()
}

func rwLock() time.Duration {
	start := time.Now()

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go rwLockRead()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go rwLockWrite()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
	return end.Sub(start)
}

func read() {
	lock.Lock()
	time.Sleep(time.Millisecond)
	lock.Unlock()
	wg.Done()
}

func write() {
	lock.Lock()
	time.Sleep(time.Millisecond)
	y = y + 1
	lock.Unlock()
	wg.Done()
}

func Lock() time.Duration {
	start := time.Now()

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()

	end := time.Now()
	return end.Sub(start)
}

func add() {
	for i := 0; i < 1000000; i++ {
		a = a + 1
	}
	wg.Done()
}

func criticalResource() {
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println("x : ", a)
}

