package main

import (
	"fmt"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	startingIndexURL := "https://www.ted.com/talks?page=2"
	aPage, _ := goquery.NewDocument(startingIndexURL)

	//linksInAPage := collectTalkLinks(aPage)
	//fmt.Println(linksInAPage)

	//lastPage(aPage)

	numBeforeNext(aPage)

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

func lastPage(doc *goquery.Document) {
	index := doc.Find(".results__pagination").Contents().Text()
	x := strings.Split(index, "|")
	fmt.Println(x)

}

func numBeforeNext(doc *goquery.Document) {

	doc.Find(".pagination__item").Each(func(i int, s *goquery.Selection) {
		num := s.Find("a").Contents().Text()
		fmt.Println(i+1, " : ", num)
	})

	//x := doc.Find(".pagination__link").Prev().Contents().Text()
	//fmt.Println(x)

}
