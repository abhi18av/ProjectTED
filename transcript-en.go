package main

import (
        "fmt"
        "github.com/PuerkitoBio/goquery"
        "strings"
)

func check(e error) {
        if e != nil {
                panic(e)
        }
}



func main() {

        url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

        doc, err := goquery.NewDocument(url)
        check(err)

        /*        topics := doc.Find(".talk-topics__list").Contents().Text()
                for _, topic := range strings.Split(topics, "  "){

                        fmt.Println(topic)
                }
        */
        langs := doc.Find(".language-list__item").Contents().Text()
        fmt.Println(langs)
        /* for _, lang := range strings.Split(langs, " "){
	                           fmt.Println(lang)
                }

        */



/*        times := doc.Find(".talk-transcript__para__time").Contents().Text()
        for _, time := range strings.Split(times, " ") {

                fmt.Println(time)
        }

        texts := doc.Find(".talk-transcript__para__text").Contents().Text()
        for _, text := range strings.Split(texts, "  ") {

                fmt.Println(text)
        }
*/}