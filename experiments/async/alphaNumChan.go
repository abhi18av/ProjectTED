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

func main() {
	//ch1 := make(chan string)
	//ch2 := make(chan string)

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	//var aleph alphanumeric

	var alephS []alphanumeric

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Execute this on 10 diferent nums
	for _, num := range numbers {
		go func(num []string, alphabets []string) {
			defer wg.Done()
			x := makeAleph(numbers, alphabets)
			alephS = append(alephS, x)
			//fmt.Println(x)
		}(numbers, alphabets)
	}

	wg.Wait()
	fmt.Println(alephS)
	//fmt.Println(alephS)
	//fmt.Println(aleph.pairAlphanumeric())
	//fmt.Println(len(aleph.anAlphabet), len(aleph.aNumber))
} // end of main()

func makeAleph(numbers []string, alphabets []string) alphanumeric {

	var aleph alphanumeric

	aleph.anAlphabet = aNum(numbers)
	aleph.aNumber = anAlph(alphabets)

	//fmt.Println(aleph.pairAlphanumeric())

	//return aleph.pairAlphanumeric()
	return aleph
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
