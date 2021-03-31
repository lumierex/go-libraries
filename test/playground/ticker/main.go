package main

import (
	"fmt"
)

func main() {
	// ticker
	//ticker := time.NewTicker(1 * time.Second)
	//cnt := 0
	//for {
	//	select {
	//	case t := <-ticker.C:
	//		{
	//			fmt.Println(t.String())
	//			cnt++
	//			if cnt == 10 {
	//				ticker.Stop()
	//			}
	//		}
	//	case <-time.After(time.Second * 2):
	//		{
	//			fmt.Println("timeout 2")
	//			break
	//		}
	//	}
	//}

	var a = [3]int{1, 2, 3}
	for k, v := range a {
		fmt.Println(k, v)
		v = 10
	}
	fmt.Println(a)
	for k, _ := range a {
		a[k] = 10
	}
	fmt.Println(a)

}
