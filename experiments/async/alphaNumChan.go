package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	numbers := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	type alphanumeric struct {
		anAlphabet string
		aNumber    string
	}

	var aleph alphanumeric

	var alephS []alphanumeric

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Execute this on 5 diferent nums
	for _, num := range numbers {
		go func(num string) {
			defer wg.Done()
			go pair1(num, timeSeed(), ch1)
			go pair2(num, timeSeed(), ch2)

			talk.first = append(talk.first, <-ch1)
			talk.second = append(talk.second, <-ch2)

		}(num)
	}

	talkS = append(talkS, talk)

	wg.Wait()
	//fmt.Println(talkS)
	fmt.Println(talk)
	//fmt.Println(len(talk.first), len(talk.second))
} // end of main()

func timeSeed() time.Duration {

	duration := time.Millisecond * time.Duration(rand.Intn(1000))
	return duration
}

func pair1(anum string, seed time.Duration, ch chan string) {

	// Sleep for arbitrary amount of time
	time.Sleep(seed)

	ch <- anum

	//return anum
}

// mechanism to select 10 alphabets at random

func tenRandomAlphabets(alphabets []string) []string {

	return tenRandomAlphabets
}
