package main

import (
	"fmt"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	startingIndexURL := "https://www.ted.com/talks?page=2"
	aPage, _ := goquery.NewDocument(startingIndexURL)

	linksInAPage := collectTalkLinks(aPage)
	fmt.Println(linksInAPage)

	lastPage := lastIndex(aPage)
	fmt.Println(lastPage)

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

func lastIndex(doc *goquery.Document) string {
	x := doc.Find(".pagination__flipper").PrevUntil(".pagination__gap").Contents().Text()
	y := strings.Split(x, "|")
	//fmt.Println(y[1])
	return y[1]
}
