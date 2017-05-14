// Have a counting of << 1 to 10 >> in each thread with arbitrary sleep time.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var threadNames = [...]string{"T1", "T2", "T3", "T4", "T5"}
	// Execute this on 5 diferent threads
	for _, thread := range threadNames {

		go func(th string) {
			oneToTen(th)
		}(thread)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func oneToTen(aThread string) {

	// Sleep for arbitrary amount of time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))

	// Line seperation for better formatting
	fmt.Println("\n\n # ", aThread, " #")

	// Prints 1 to 10 on the screen
	fmt.Println(">>>>>>>>>>>>>>>")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("<<<<<<<<<<<<<<<")
}
