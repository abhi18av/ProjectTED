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

// this should return an array of strings => ["langs"]
func langs(doc *goquery.Document) []string {

	var langsList []string

	langs := doc.Find(".talk-transcript__language").Contents().Text()

	//	fmt.Println(langs)
	langsSeparated := strings.Split(langs, "\n")

	for i := 1; i < len(langsSeparated)-1; i++ {
		//fmt.Println(i, ":", langsSeparated[i])
		langsList = append(langsList, langsSeparated[i])
	}

	return langsList
}

func main() {

	//url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"
	url := "https://www.ted.com/talks/jill_bolte_taylor_s_powerful_stroke_of_insight/transcript?language=en"
	//url := "https://www.ted.com/talks/tony_robbins_asks_why_we_do_what_we_do/transcript?language=en"
	//url := "https://www.ted.com/talks/shawn_achor_the_happy_secret_to_better_work/transcript?language=en"
	//url := "https://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action/transcript?language=en"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(langs(doc)))
}
