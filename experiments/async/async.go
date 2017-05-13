package main

import (
	"fmt"
	"strings"

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

		go talk_texts(url, textChannel)
	}

	/*
		for _, url := range urls {

			go fetch(url, ch)
		}

	*/

	for range urls {
		fmt.Println(<-textChannel)
	}

}

/*
func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		println("Error with the fetch")
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()
	if err != nil {
		println("Error with IO")

	}

	//	ch <- fmt.Sprintf("%7d    %s", nbytes, url)

	ch <- fmt.Println(resp.Body)
}
*/

func talk_texts(url string, textChannel chan<- []string) {
	page, _ := goquery.NewDocument(url)

	texts := page.Find(".talk-transcript__para__text").Contents().Text()
	var para []string
	for _, text := range strings.Split(texts, "\n\n") {

		//fmt.Println(text)
		para = append(para, text)
	}

	textChannel <- para
}
