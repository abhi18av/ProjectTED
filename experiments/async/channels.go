package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(timeSeed())
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(timeSeed())
		}
	}()

	msg1 := <-c1
	fmt.Println(msg1)
	msg2 := <-c2
	fmt.Println(msg2)
}

func timeSeed() time.Duration {

	duration := time.Millisecond * time.Duration(rand.Intn(1000))
	return duration
}
