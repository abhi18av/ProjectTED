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
		title string
	}

	var talk collections

	// Execute this on 5 diferent threads
	for _, thread := range threadNames {
		go func(thread string) {
			defer wg.Done()
			talk.title = currentThread(thread)
		}(thread)
	}

	wg.Wait()
} // end of main()

func currentThread(aThread string) string {

	// Sleep for arbitrary amount of time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))

	// Line seperation for better formatting
	fmt.Println("# ", aThread, " #")
	return (aThread)
	/*
		// Prints 1 to 10 on the screen
		fmt.Println(">>>>>>>>>>>>>>>")
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
		fmt.Println("<<<<<<<<<<<<<<<")
	*/
}
