package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	startingIndexURL := "https://www.ted.com/talks?page=1"
	aPage, _ := goquery.NewDocument(startingIndexURL)
	collectTalkLinks(aPage)
}

func collectTalkLinks(doc *goquery.Document) {
	urls, _ := doc.Find(".ga-link").Attr("href")
	fmt.Println(urls)
	//return urls
}
