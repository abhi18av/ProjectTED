package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func posted(doc *goquery.Document) {
	posted := doc.Find(".meta__item").Contents().Text()
	fmt.Println(posted)
}

func rated(doc *goquery.Document) {
	rated := doc.Find(".meta__row").Contents().Text()
	fmt.Println(rated)
}

func title(doc *goquery.Document) {
	title := doc.Find(".m5").Contents().Text()
	fmt.Println(strings.Split(title, "\n")[2])
}

func topics(doc *goquery.Document) {
	topics := doc.Find(".talk-topics__list").Contents().Text()
	for _, topic := range strings.Split(topics, "  ") {

		fmt.Println(topic)
	}
}

func langs(doc *goquery.Document) {

	// TODO: Need a table for lang codes from
	// https://www.ted.com/participate/translate/our-languages
	langs := doc.Find(".talk-transcript__language").Contents().Text()
	fmt.Println(langs)
}

func times(doc *goquery.Document) {
	times := doc.Find(".talk-transcript__para__time").Contents().Text()
	for _, time := range strings.Split(times, " ") {

		fmt.Println(time)
	}
}

func texts(doc *goquery.Document) {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	for _, text := range strings.Split(texts, "  ") {

		fmt.Println(text)
	}
}

func main() {

	url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	posted(doc)
	rated(doc)
	title(doc)
	topics(doc)
	langs(doc)

	times(doc)
	texts(doc)
}
