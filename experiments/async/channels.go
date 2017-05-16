package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	type structOfTwo struct {
		first string
		two   string
	}

	var two structOfTwo

	go func1(c1)
	two.first = <-c1
	//fmt.Println(msg1)

	go func2(c2)
	two.two = <-c2
	//fmt.Println(msg2)

	fmt.Println(two)

}

func timeSeed() time.Duration {

	duration := time.Millisecond * time.Duration(rand.Intn(100000))
	return duration
}

func func1(ch chan string) {
	ch <- "1"
	time.Sleep(timeSeed())

}

func func2(ch chan string) {
	ch <- "2"
	time.Sleep(timeSeed())

}
