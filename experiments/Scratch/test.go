package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	alphabets := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var arr []string

	for _, x := range alphabets {

		if len(arr) == 10 {
			break
		}

		rand.Seed(time.Time.UnixNano(time.Now()))

		if rand.Intn(10)%2 == 0 {

			arr = append(arr, x)
		}
	}

	fmt.Println(arr)
}
