package main

import "fmt"

func main() {

	//var a[len] type
	//var a [5]int
	//a[0] = 1

	var a = [5]int{1, 2, 3, 4, 5}
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	// 索引初始化
	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Println(len(a), a)
	fmt.Println(len(a1), a1)
	fmt.Print(len(str), str)
	//
	//
	var a2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(a2)

	changeArray(a)
	fmt.Println(a)

	changeArrayWithPointer(&a)
	fmt.Println(a)

}

func changeArray(a [5]int) {
	a[0] = 10
	for k, v := range a {
		fmt.Println(k, v)
	}
}

func changeArrayWithPointer(a *[5]int) {
	a[0] = 10
	for k, v := range a {
		fmt.Println(k, v)
	}
}
