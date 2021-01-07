package groutine

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

var m1 sync.Map

func set(key string, value int) {
	m[key] = value
}

func get(key string) int {
	//if i, ok := m[key]; ok {
	//	return i
	//}
	return m[key]
}

func mapConcurrent() {
	wg := sync.WaitGroup{}
	for i := 0; i < 4000; i++ {
		wg.Add(1)
		go func(val int) {
			key := strconv.Itoa(val)
			set(key, val)
			fmt.Println("key", key, "value", get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func mapConcurrentSave() {

	wg := sync.WaitGroup{}
	for i := 0; i < 4000; i++ {
		wg.Add(1)
		go func(val int) {
			key := strconv.Itoa(val)
			//fmt.Println("start key", key, i)
			m1.Store(key, val)
			load, ok := m1.Load(key)
			if ok {
				fmt.Println("key", key, "value", load)
			}
			wg.Done()
		}(i)
	}
}
