package main

import (
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	urls := []string{
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en",
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=zh-cn",
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=de",
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=ru",
	}

	textChannel := make(chan []string)

	for _, url := range urls {

		go talkTexts(url, textChannel)
	}

	//var texts []string
	for range urls {
		//texts = append(texts, <-textChannel)
		fmt.Println(<-textChannel)
	}

	//fmt.Println(texts)
}

func talkTexts(url string, textChannel chan<- []string) {
	page, _ := goquery.NewDocument(url)

	texts := page.Find(".talk-transcript__para__text").Contents().Text()
	var para []string
	for _, text := range strings.Split(texts, "\n\n") {

		//fmt.Println(text)
		para = append(para, text)

	}
	textChannel <- para
}
