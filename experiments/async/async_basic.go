// Have a counting of << 1 to 10 >> in each thread with arbitrary sleep time.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var threadNames = [...]string{"T1", "T2", "T3", "T4", "T5"}

	wg.Add(len(threadNames))

	type collections struct {
		thread []string
		count  [][]int
	}

	var talk collections

	// Execute this on 5 diferent threads
	for _, thread := range threadNames {
		go func(thread string) {
			defer wg.Done()
			talk.thread = append(talk.thread, currentThread(thread))
			talk.count = append(talk.count, one2ten())
		}(thread)
	}

	wg.Wait()

	fmt.Println(talk)
} // end of main()

func currentThread(aThread string) string {

	// Sleep for arbitrary amount of time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10000)))

	// Line seperation for better formatting
	return aThread
}

func one2ten() []int {

	var arr []int

	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
	}

	return arr

}
