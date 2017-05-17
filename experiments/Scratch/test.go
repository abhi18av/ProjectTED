package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type alphanumeric struct {
	anAlphabet string
}

func main() {

	var wg sync.WaitGroup
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	wg.Add(10)
	go func() {
		defer wg.Done()
		makeAleph(alphabets)
	}()
	//fmt.Println(makeAleph(alphabets))

	wg.Wait()
}

func makeAleph(alphabets []string) {

	var chanAlphabet chan string

	go aNum(chanAlphabet, alphabets)

	x := <-chanAlphabet

	fmt.Println(x)

}

func randomIndex() int {
	randTime := time.Time.UnixNano(time.Now())

	rand.Seed(randTime)

	return rand.Intn(10)
}

func aNum(chanNumber chan string, numbers []string) {

	chanNumber <- numbers[randomIndex()]

}
