package main

import (
	"fmt"
)

func countTill(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func main() {

	fmt.Println("Please enter the num: ")
	var x int
	fmt.Scanf("%d", &x)
	println("\nBeginning the count: ")
	// countTill(x)

}
