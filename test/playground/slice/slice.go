package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SDuck struct {
	id   int64
	name string
}

func main() {

	var a []int
	a = append(a, 1)
	fmt.Println(a, cap(a), a == nil)

	b := make([]int, 0)
	fmt.Println(b, cap(b), b == nil)

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)

	// slice substring
	hstr := "hello world"
	i := strings.Index(hstr, "w")
	s := hstr[i:]
	fmt.Println(s)
	fmt.Println(hstr[:])
	s4 := hstr[0:strings.Index(hstr, " ")]
	fmt.Println(s4)

	a1 := 10

	passWithValue(a1)
	fmt.Println(a1)
	b1 := 10
	passWithPointer(&b1)
	fmt.Println(b1)

	fmt.Println("=======================")

	config := make(map[string]interface{})

	config["port"] = 9000
	config["host"] = "127.0.0.1"

	fmt.Println(config)
	for k, v := range config {
		fmt.Println(k, v)
	}
	x, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(x))
	password, ok := config["password"]
	if !ok {
		fmt.Println("password does not exist", password)
	}

}
func passWithValue(a int) {
	a = 10
}

func passWithPointer(a *int) {
	*a = 11
}
