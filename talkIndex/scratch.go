package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://blog.golang.org")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("h3 a").Attr("href")
		fmt.Println(i+1, link)
	})
}
