package main

import (
	"fmt"
	"math/rand"
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

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	//var aleph alphanumeric
	//var alephS []alphanumeric

	x := makeAleph(numbers, alphabets)
	fmt.Println(x)
}

func makeAleph(numbers []string, alphabets []string) string {

	var chanAlphabet chan string
	var chanNumber chan string

	go aNum(chanNumber, numbers)
	go anAlph(chanAlphabet, alphabets)

	var aleph alphanumeric

	aleph.anAlphabet = <-chanAlphabet
	aleph.aNumber = <-chanNumber

	return aleph.pairAlphanumeric()
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
