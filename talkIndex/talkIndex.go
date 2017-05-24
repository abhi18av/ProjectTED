package main

import (
	"github.com/PuerkitoBio/goquery"
)

func main() {

	startingIndexURL := "https://www.ted.com/talks?page=2"
	aPage, _ := goquery.NewDocument(startingIndexURL)
	collectTalkLinks(aPage)
}

func collectTalkLinks(doc *goquery.Document) []string {
	var urls []string

	doc.Find(".h9").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		//fmt.Println(i+1, link)
		urls = append(urls, link)
	})

	//fmt.Println(urls)
	return urls
}
