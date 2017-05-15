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
		title [][]int
	}

	var talk collections

	// Execute this on 5 diferent threads
	for _, thread := range threadNames {
		go func(thread string) {
			defer wg.Done()
			talk.title = append(talk.title, currentThread(thread))
		}(thread)
	}

	wg.Wait()

	fmt.Println(talk)
} // end of main()

func currentThread(aThread string) []int {

	// Sleep for arbitrary amount of time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10000)))

	// Line seperation for better formatting
	// fmt.Println("# ", aThread, " #")

	var one2ten []int

	// Prints 1 to 10 on the screen
	//fmt.Println(">>>>>>>>>>>>>>>")
	for i := 1; i <= 10; i++ {
		one2ten = append(one2ten, i)
	}
	//fmt.Println("<<<<<<<<<<<<<<<")

	return one2ten
}
