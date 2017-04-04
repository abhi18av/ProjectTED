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

func main() {

	url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)
	check(err)

	posted := doc.Find(".meta__item").Contents().Text()
	fmt.Println(posted)

	rated := doc.Find(".meta__row").Contents().Text()
	fmt.Println(rated)

	title := doc.Find(".m5").Contents().Text()
	fmt.Println(strings.Split(title, "\n")[2])

	topics := doc.Find(".talk-topics__list").Contents().Text()
	for _, topic := range strings.Split(topics, "  ") {

		fmt.Println(topic)
	}

	// TODO: Need a table for lang codes from
	// https://www.ted.com/participate/translate/our-languages
	langs := doc.Find(".talk-transcript__language").Contents().Text()
	fmt.Println(langs)

	times := doc.Find(".talk-transcript__para__time").Contents().Text()
	for _, time := range strings.Split(times, " ") {

		fmt.Println(time)
	}

	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	for _, text := range strings.Split(texts, "  ") {

		fmt.Println(text)
	}

}
