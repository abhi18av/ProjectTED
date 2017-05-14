// Have a counting of << 1 to 10 >> in each thread with arbitrary sleep time.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
	"German":              "de",
	"Russian":             "ru",
}

func main() {
	var wg sync.WaitGroup

	var threadNames = [...]string{"T1", "T2", "T3", "T4", "T5"}

	wg.Add(len(threadNames))
	baseURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language="

	//	var urls []string

	for _, value := range langCodes {
		x := baseURL + value
		fmt.Println(x)
	}
	// Execute this on 5 diferent threads
	for _, thread := range threadNames {
		go func(th string) {
			defer wg.Done()
			currentThread(th)
		}(thread)
	}

	wg.Wait()
}

func currentThread(aThread string) {

	// Sleep for arbitrary amount of time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))

	// Line seperation for better formatting
	fmt.Println("# ", aThread, " #")

	/*
		// Prints 1 to 10 on the screen
		fmt.Println(">>>>>>>>>>>>>>>")
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
		fmt.Println("<<<<<<<<<<<<<<<")
	*/
}
