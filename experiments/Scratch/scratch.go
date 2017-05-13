package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		x := sqr(i)
		fmt.Println(i, " : ", x)
	}

}

func sqr(x int) int {
	return x * x
}
