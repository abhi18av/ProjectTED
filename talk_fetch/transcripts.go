package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func posted(doc *goquery.Document) string {
	posted := doc.Find(".meta__item").Contents().Text()
	p := strings.Split(posted, "\n")
	//fmt.Println(p[3])
	return (p[3])

}

func rated(doc *goquery.Document) {

	rated := doc.Find(".meta__row").Contents().Text()

	r := strings.Split(rated, "\n")
	fmt.Println(r[3])

	/*
	   rx := strings.Split(r[3], ",")

	   	for _, x := range rx{
	   		append(ls,x)
	   	}
	*/

	//println(len(rx))
	//println(r[0])
	//println(r[1])
	//return(p[3])
}

func local_title(doc *goquery.Document) {
	title := doc.Find(".m5").Contents().Text()
	fmt.Println(strings.Split(title, "\n")[2])
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
	//http://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en
	//http://www.ted.com/talks/jill_bolte_taylor_s_powerful_stroke_of_insight/transcript?language=en
	//http://www.ted.com/talks/tony_robbins_asks_why_we_do_what_we_do/transcript?language=en
	//http://www.ted.com/talks/shawn_achor_the_happy_secret_to_better_work/transcript?language=en
	//http://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action/transcript?language=en

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	//	posted(doc)
	//	rated(doc)
	//	local_title(doc)

	langs(doc)
	//	times(doc)
	//	texts(doc)

}
