package main

import "fmt"

type alphanumeric struct {
	anAlphabet string
	aNumber    string
}

func main() {

	var aleph alphanumeric
	var alephS []alphanumeric

	aleph.anAlphabet = "a"
	aleph.aNumber = "1"

	fmt.Println(aleph.pairAlphanumeric())
}

func (someStruct alphanumeric) pairAlphanumeric() string {

	return someStruct.aNumber + someStruct.anAlphabet

}
