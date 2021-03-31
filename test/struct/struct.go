package main

import (
	"fmt"
	"strings"
)

type Hello struct {
	Name string
}

var m = make(map[string]Hello)

func structIsEqual() {

}

func main() {
	splits := strings.Split("xxx", ",")
	fmt.Println(splits)
}
