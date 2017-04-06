package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func posted(doc *goquery.Document) {

	posted := doc.Find(".meta__item").Contents().Text()
	fmt.Println(posted)

}

func topics(doc *goquery.Document) {
	topics := doc.Find(".talk-topics__list").Contents().Text()
	for _, topic := range strings.Split(topics, "  ") {

		fmt.Println(topic)
	}
}
func main() {

	url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"
	doc, err := goquery.NewDocument(url)
	check(err)
	posted(doc)
	topics(doc)
}
