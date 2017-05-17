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
	n := 10 // number of codes you want to print

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(numbers []string, alphabets []string) {
			defer wg.Done()
			makeAleph(numbers, alphabets)
		}(numbers, alphabets)
	}

	wg.Wait()
} // end of main()

func makeAleph(numbers []string, alphabets []string) {

	var aleph alphanumeric

	aleph.anAlphabet = aNum(numbers)
	aleph.aNumber = anAlph(alphabets)

	fmt.Println(aleph.pairAlphanumeric())

	//return aleph.pairAlphanumeric()
}

func randomIndex() int {
	randTime := time.Time.UnixNano(time.Now())

	rand.Seed(randTime)

	return rand.Intn(10)
}

func aNum(numbers []string) string {

	return numbers[randomIndex()]

}

func anAlph(alphabets []string) string {

	return alphabets[randomIndex()]

}
