package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func1(c1)
	go func2(c2)

	var msg1 string
	var msg2 string

	select {
	case msg1 = <-c1:
		fmt.Println(msg1)
	case msg2 = <-c2:
		fmt.Println(msg2)
	}

	fmt.Println("Before Exit")
	fmt.Println(msg1, msg2)
}

func timeSeed() time.Duration {

	duration := time.Millisecond * time.Duration(rand.Intn(10000))
	return duration
}

func func1(ch chan string) {
	ch <- "from 1"
	time.Sleep(timeSeed())

}

func func2(ch chan string) {
	ch <- "from 2"
	time.Sleep(timeSeed())

}
