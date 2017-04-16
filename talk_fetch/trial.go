package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
	"German":              "de",
	"Russian":             "ru",
}

func texts(doc *goquery.Document) {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	for _, text := range strings.Split(texts, "  ") {

		fmt.Println(text)
	}
}

// this should return an array of strings => ["langs"]
func langs(doc *goquery.Document) []string {

	var langsList []string

	langs := doc.Find(".talk-transcript__language").Contents().Text()

	//fmt.Println(langs)
	langsSeparated := strings.Split(langs, "\n")

	for i := 1; i < len(langsSeparated)-1; i++ {
		//fmt.Println(i, ":", langsSeparated[i])
		langsList = append(langsList, langsSeparated[i])
	}

	return langsList
}

func main() {

	var doc []*goquery.Document

	url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	page, _ := goquery.NewDocument(url)
	doc = append(doc, page)

	fmt.Println(doc[0])

	//fmt.Println(langs(doc)[1])

	/*
			var availableLangs []string

			for _, code := range langs(doc[0]) {

				if langCodes[code] == "" {
					//fmt.Println(code, " : Not available")
				} else {
					fmt.Println(langCodes[code])
					availableLangs = append(availableLangs, langCodes[code])
				}
			}

		fmt.Println(availableLangs)
	*/

	//newURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=" + availableLangs[0]
	//doc[1], _ = goquery.NewDocument(newURL)

	//texts(doc[1])

	//texts(goquery.NewDocument(newURL))
}
