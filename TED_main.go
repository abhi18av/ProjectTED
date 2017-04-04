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

	url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)
	check(err)

	//speaker := doc.Find(".talk-speaker__name").Contents().Text()
	//fmt.Println(speaker)

	//title := doc.Find(".player-hero__title__content").Contents().Text()
	//fmt.Println(title)

	duration := doc.Find(".player-hero__meta").Contents().Text()
	//fmt.Println(duration)

	/*
		for _, x := range strings.Split(duration, "\n") {
			fmt.Println(x)
			println("~~~~~~")
		}
	*/
	x := strings.Split(duration, "\n")
	fmt.Println(x[6])

	/*
		time_filmed := doc.Find(".player-hero__meta__label").Contents().Text()

		fmt.Println(time_filmed)

		talk_views_count := doc.Find("#sharing-count").Contents().Text()
		fmt.Println(talk_views_count)

		talk_topics := doc.Find(".talk-topics__list").Contents().Text()
		fmt.Println(talk_topics)

		talk_comments_count := doc.Find(".h11").Contents().Text()
		fmt.Println(talk_comments_count)

	*/
}
