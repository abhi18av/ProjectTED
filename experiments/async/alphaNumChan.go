package main

import (
	"math/rand"
	"sync"
	"time"
)

type alphanumeric struct {
	anAlphabet string
	aNumber    string
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var aleph alphanumeric

	var alephS []alphanumeric

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Execute this on 10 diferent nums
	for _, num := range numbers {
		go func(num string) {
			defer wg.Done()
			tenAlphabets := tenRandomAlphabets(alphabets)

			pairAlphanumeric(num, tenAlphabets, ch1, ch2)

			alephS.anAlphabet = <-ch1
			alephS.aNumber = <-ch2

		}(num)
	}

	alephS = append(alephS, aleph)

	wg.Wait()

	//fmt.Println(alephS)
	//fmt.Println(aleph.pairAlphanumeric())
	//fmt.Println(len(aleph.anAlphabet), len(aleph.aNumber))
} // end of main()

func timeSeed() time.Duration {

	duration := time.Millisecond * time.Duration(rand.Intn(1000))
	return duration
}

func (someStruct alphanumeric) pairAlphanumeric() string {

	return someStruct.aNumber + someStruct.anAlphabet

}

// mechanism to select 10 alphabets at random
func tenRandomAlphabets(alphabets []string) []string {

	var arr []string

	for _, x := range alphabets {

		if len(arr) == 10 {
			break
		}

		randTime := time.Time.UnixNano(time.Now())
		rand.Seed(randTime)

		if rand.Int63n(randTime)%2 == 0 {

			arr = append(arr, x)
		}
	}
	return arr
}

func makeAleph(numbers []string, alphabets [] ) {

}
