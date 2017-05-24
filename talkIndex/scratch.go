package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, _ := goquery.NewDocument("https://blog.golang.org")

	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("h3 a").Attr("href")
		fmt.Println(i+1, link)
	})
}
