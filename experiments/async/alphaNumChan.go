package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type alphanumeric struct {
	anAlphabet string `json:"anAlphabet"`
	aNumber    string `json:"aNumber"`
}

func (someStruct alphanumeric) pairAlphanumeric() string {

	return someStruct.aNumber + someStruct.anAlphabet

}

func main() {
	//ch1 := make(chan string)
	//ch2 := make(chan string)

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	//var aleph alphanumeric

	var alephS []alphanumeric

	var wg sync.WaitGroup

	n := 10
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(numbers []string, alphabets []string) {
			defer wg.Done()
			x := makeAleph(numbers, alphabets)
			//fmt.Println(x.anAlphabet)
			alephS = append(alephS, x)
			//fmt.Println(x)
		}(numbers, alphabets)
	}

	wg.Wait()

	fmt.Println(alephS[0])

	body, err := json.Marshal(alephS[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

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
