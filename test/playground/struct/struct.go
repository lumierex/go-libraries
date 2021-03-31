package main

import (
	"encoding/json"
	"fmt"
)

type Duck interface {
	MakeNoise()
}

type Chicken struct {
	id   int64
	name string
}

func main() {
	chicken := Chicken{
		id:   12,
		name: "hi",
	}
	chickenJson, err := json.Marshal(chicken)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(chickenJson))

}
