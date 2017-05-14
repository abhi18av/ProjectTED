package main

import (
	"fmt"
	"strings"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
	"German":              "de",
	"Russian":             "ru",
}

func main() {

	var wg sync.WaitGroup

	baseURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	var urls []string

	for _, value := range langCodes {
		newURL := baseURL + "/transcript?language=" + value
		//fmt.Println(x)
		urls = append(urls, newURL)
	}
	fmt.Println(urls)
	wg.Add(len(urls))

	for _, url := range urls {

		go func(url string) {
			fmt.Println(talkTexts(url))
			defer wg.Done()
		}(url)
	}

	//var texts []string

	//fmt.Println(texts)
	wg.Wait()
}

func talkTexts(url string) []string {
	//textChannel := make(chan []string)
	var para []string
	page, _ := goquery.NewDocument(url)

	texts := page.Find(".talk-transcript__para__text").Contents().Text()

	for _, text := range strings.Split(texts, "\n\n") {

		//fmt.Println(text)
		para = append(para, text)

	}
	return para
}
