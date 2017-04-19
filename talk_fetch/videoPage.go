package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var langCodes = map[string]string{

	"English": "en",
	"French":  "fr",
}

func AvailableSubtitles(doc *goquery.Document) int64 {

	subtitles := doc.Find(".player-hero__meta__link").Contents().Text()
	//fmt.Println(subtitles)

	//for _, x := range strings.Split(subtitles, "\n") {
	//fmt.Println(x)
	//println("~~~~~~")
	//}

	y := strings.Split(subtitles, "\n")
	z := strings.Split(y[3], " ")[0]
	numOfSubtitles, _ := strconv.ParseInt(z, 10, 32)
	return numOfSubtitles
}

func Speaker(doc *goquery.Document) string {
	speaker := doc.Find(".talk-speaker__name").Contents().Text()
	//fmt.Println(speaker)
	speaker = strings.Trim(speaker, "\n")
	return speaker
}

/*
// This is now taken from the transcripts page
func title(doc *goquery.Document) {
	title := doc.Find(".player-hero__title__content").Contents().Text()
	fmt.Println(title)
}
*/

func Duration(doc *goquery.Document) string {

	duration := doc.Find(".player-hero__meta").Contents().Text()
	//fmt.Println(duration)

	//for _, x := range strings.Split(duration, "\n") {
	//	fmt.Println(x)
	//	println("~~~~~~")
	//}

	x := strings.Split(duration, "\n")
	fmt.Println(x[6])
	return x[6]

}

// TimeFilmed : Time at which the talk was filmed
func TimeFilmed(doc *goquery.Document) string {

	time_filmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(time_filmed)

	y := strings.Split(time_filmed, "\n")
	//fmt.Println(y[11])
	return y[11]
}

func TalkViewsCount(doc *goquery.Document) string {

	talk_views_count := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talk_views_count)

	a := strings.Split(talk_views_count, "\n")
	b := strings.TrimSpace(a[2])
	//fmt.Println(b)
	return b

}

func TalkTopicsList(doc *goquery.Document) []string {

	talk_topics := doc.Find(".talk-topics__list").Contents().Text()

	c := strings.Split(talk_topics, "\n")
	var topics []string
	for i := 3; i < len(c); i++ {
		//fmt.Println(c[i])
		if c[i] == "" {

		} else {
			topics = append(topics, c[i])
		}
	}
	return topics
}

func TalkCommentsCount(doc *goquery.Document) string {

	talk_comments_count := doc.Find(".h11").Contents().Text()
	//fmt.Println(talk_comments_count)
	d := strings.Split(talk_comments_count, " ")
	//fmt.Println(d[0])
	return strings.TrimLeft(d[0], "\n")
}

func main() {

	//url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"
	//url := "https://www.ted.com/talks/jill_bolte_taylor_s_powerful_stroke_of_insight"
	url := "https://www.ted.com/talks/tony_robbins_asks_why_we_do_what_we_do"
	//url := "https://www.ted.com/talks/shawn_achor_the_happy_secret_to_better_work"
	//url := "https://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action"

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	//speaker(doc)

	//println(availableSubtitles(doc))

	/*
		// This is now taken from the transcripts pag
		//title(doc)
	*/

	duration(doc)
	//time_filmed(doc)
	//talk_views_count(doc)
	//talk_topics_list(doc)
	//talk_comments_count(doc)

}
