package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	p := person{name: "bond", age: 34}
	fmt.Println(p)
}
