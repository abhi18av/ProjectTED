// Have a counting of << 1 to 10 >> in each thread.
package main

import "fmt"

func main() {

	var threadNames = [...]string{"T1", "T2", "T3", "T4", "T5"}
	// Execute this on 5 diference threads
	for _, thread := range threadNames {
		oneToTen(thread)
	}
}

func oneToTen(channelName string) {

	fmt.Println("\n\n # ", channelName, " #")

	fmt.Println(">>>>>>>>>>>>>>>")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("<<<<<<<<<<<<<<<")
}
