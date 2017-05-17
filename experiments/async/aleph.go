package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type alphanumeric struct {
	anAlphabet string
	aNumber    string
}

func (someStruct alphanumeric) pairAlphanumeric() string {

	return someStruct.aNumber + someStruct.anAlphabet

}

func main() {

	var wg sync.WaitGroup

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	//var aleph alphanumeric
	//var alephS []alphanumeric

	wg.Add(10)
	go func(numbers []string, alphabets []string) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			makeAleph(numbers, alphabets)
		}
	}(numbers, alphabets)
	wg.Wait()
} // end of main()

func makeAleph(numbers []string, alphabets []string) {

	var chanAlphabet chan string
	var chanNumber chan string
	var wg sync.WaitGroup

	wg.Add(10)

	go func() {
		defer wg.Done()
		aNum(chanNumber, numbers)
	}()

	go func() {
		defer wg.Done()
		anAlph(chanAlphabet, alphabets)
	}()

	var aleph alphanumeric

	aleph.anAlphabet = <-chanAlphabet
	aleph.aNumber = <-chanNumber

	fmt.Println(aleph.pairAlphanumeric())
	//return aleph.pairAlphanumeric()
	wg.Wait()
}

func randomIndex() int {
	randTime := time.Time.UnixNano(time.Now())

	rand.Seed(randTime)

	return rand.Intn(10)
}

func aNum(chanNumber chan string, numbers []string) {

	chanNumber <- numbers[randomIndex()]

}

func anAlph(chanAlphabet chan string, alphabets []string) {

	chanAlphabet <- alphabets[randomIndex()]

}
