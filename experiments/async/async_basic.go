package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var threadNames = [...]string{"T1", "T2", "T3", "T4", "T5", "T6", "T7", "T8", "T9", "T10"}

	wg.Add(len(threadNames))

	type collections struct {
		first  []string
		second []string
	}

	var talk collections

	// Execute this on 5 diferent threads
	for _, thread := range threadNames {
		go func(thread string) {
			defer wg.Done()
			talk.first = append(talk.first, pair1(thread, timeSeed()))
			talk.second = append(talk.second, pair2(thread, timeSeed()))
		}(thread)
	}

	wg.Wait()

	fmt.Println(talk)
	fmt.Println(len(talk.first), len(talk.second))
} // end of main()

func timeSeed() time.Duration {

	duration := time.Millisecond * time.Duration(rand.Intn(1000))
	return duration
}
func pair1(aThread string, seed time.Duration) string {

	// Sleep for arbitrary amount of time
	time.Sleep(seed)

	return aThread
}

func pair2(aThread string, seed time.Duration) string {

	// Sleep for arbitrary amount of time
	time.Sleep(seed)

	return aThread
}
