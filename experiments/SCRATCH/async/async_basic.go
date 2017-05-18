package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var threadNames = [...]string{"T1", "T2", "T3", "T4", "T5", "T6", "T7", "T8", "T9", "T10"}

	var wg sync.WaitGroup
	wg.Add(len(threadNames))

	type collections struct {
		first  []string
		second []string
	}

	var talk collections

	var talkS []collections

	c1 := make(chan string)
	c2 := make(chan string)

	// Execute this on 5 diferent threads
	for _, thread := range threadNames {
		go func(thread string) {
			defer wg.Done()
			go pair1(thread, timeSeed(), c1)
			go pair2(thread, timeSeed(), c2)

			talk.first = append(talk.first, <-c1)
			talk.second = append(talk.second, <-c2)

		}(thread)
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

func pair1(aThread string, seed time.Duration, ch chan string) {

	// Sleep for arbitrary amount of time
	time.Sleep(seed)

	ch <- aThread

	//return aThread
}

func pair2(aThread string, seed time.Duration, ch chan string) {

	// Sleep for arbitrary amount of time
	time.Sleep(seed)

	ch <- aThread

	//return aThread
}
