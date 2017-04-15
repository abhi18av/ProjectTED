package main

import "fmt"

var lang_codes = map[string]string{

	"English": "en",
	"French":  "fr",
}

func main() {
	fmt.Println(lang_codes["French"])
}
