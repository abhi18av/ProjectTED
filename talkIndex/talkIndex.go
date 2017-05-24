package main

import (
	"strings"

	"fmt"
	"sync"

	"strconv"

	"encoding/json"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	startingIndexURL := "https://www.ted.com/talks?page=1"
	page, _ := goquery.NewDocument(startingIndexURL)

	lastIndex := lastIndex(page)
	numLastIndex, _ := strconv.ParseInt(lastIndex, 10, 64)
	fmt.Println(lastIndex)

	//linksInAPage := collectTalkLinks(page)
	//fmt.Println(linksInAPage)

	mainTedURL := "https://www.ted.com"

	var allTalksLinks [][]string
	var wg sync.WaitGroup
	wg.Add(int(numLastIndex))
	for _, index := range lastIndex {
		URL := mainTedURL + "/talks?page=" + string(index)

		go func(URL string) {
			defer wg.Done()
			aPage, _ := goquery.NewDocument(URL)
			links := collectTalkLinks(aPage)

			allTalksLinks = append(allTalksLinks, links)
		}(URL)

	}

	wg.Wait()

	//fmt.Println(len(allTalksLinks))
	x, _ := json.Marshal(collectTalkLinks)
	fmt.Println(x)
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
