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

	speaker := doc.Find(".talk-speaker__name").Contents().Text()
	fmt.Println(speaker)

	title := doc.Find(".player-hero__title__content").Contents().Text()
	fmt.Println(title)

	duration := doc.Find(".player-hero__meta").Contents().Text()
	//fmt.Println(duration)

	//for _, x := range strings.Split(duration, "\n") {
	//	fmt.Println(x)
	//	println("~~~~~~")
	//}

	x := strings.Split(duration, "\n")
	fmt.Println(x[6])

	time_filmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(time_filmed)

	y := strings.Split(time_filmed, "\n")
	fmt.Println(y[11])

	talk_views_count := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talk_views_count)

	a := strings.Split(talk_views_count, "\n")
	b := strings.TrimSpace(a[2])
	fmt.Println(b)

	talk_topics := doc.Find(".talk-topics__list").Contents().Text()
	//fmt.Println(talk_topics)

func topics(doc *goquery.Document) {
	topics := doc.Find(".talk-topics__list").Contents().Text()
	for _, topic := range strings.Split(topics, "  ") {

		fmt.Println(topic)
	}
}



	c := strings.Split(talk_topics, "\n")

	//	fmt.Println(c[3:len(c)])

	for i := 3; i < len(c); i++ {
		fmt.Println(c[i])
	}

	talk_comments_count := doc.Find(".h11").Contents().Text()
	//fmt.Println(talk_comments_count)
	d := strings.Split(talk_comments_count, " ")
	fmt.Println(d[0])

}
