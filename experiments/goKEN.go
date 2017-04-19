package main

import (
	//"encoding/json"
	"fmt"
	"strings"
	//"github.com/huandu/xstrings""
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	//"os"
	"strconv"
	"github.com/Jeffail/gabs"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
	"German":              "de",
	"Russian":             "ru",
}

type kenJSON struct {
	// The link to the main video
	TalkLink string `json:"talk_link"`

	AvailableSubtitlesCount string

	Speaker string

	Duration string

	TimeFilmed string

	TalkViewsCount string

	TalkTopicsList string

	TalkCommentsCount string


/*
	// The main container for the Transcript text
	Transcript struct {
		En   []string `json:"en"`
		ZhCn []string `json:"zh-cn"`
		De   []string `json:"de"`
		Ru   []string `json:"ru"`
	} `json:"transcript"`
*/
}

// From the main page or the English language transcript
func title(doc *goquery.Document) string {
	title := doc.Find(".player-hero__title__content").Contents().Text()
	//fmt.Println(title)
	return title
}

func texts(doc *goquery.Document) []string{
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	var paragraphs []string
	for _, text := range strings.Split(texts, "  ") {

		//fmt.Println(text)
		paragraphs = append(paragraphs, text)
	}
	return paragraphs
}



func availableSubtitles(doc *goquery.Document) int64 {

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

func speaker(doc *goquery.Document) string{
	speaker := doc.Find(".talk-speaker__name").Contents().Text()
	//fmt.Println(speaker)
	speaker = strings.Trim(speaker, "\n")
	return speaker
}



func duration(doc *goquery.Document) string {

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

func time_filmed(doc *goquery.Document) string{

	time_filmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(time_filmed)

	y := strings.Split(time_filmed, "\n")
	//fmt.Println(y[11])
	return y[11]
}

func talk_views_count(doc *goquery.Document) string{

	talk_views_count := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talk_views_count)

	a := strings.Split(talk_views_count, "\n")
	b := strings.TrimSpace(a[2])
	//fmt.Println(b)
	return b

}

func talk_topics_list(doc *goquery.Document) []string{

	talk_topics := doc.Find(".talk-topics__list").Contents().Text()

	c := strings.Split(talk_topics, "\n")
	var topics []string
	for i := 3; i < len(c); i++ {
		//fmt.Println(c[i])
		if c[i] ==""{

		}else
		{
		topics = append(topics, c[i])
	}
	}
	return topics 
}

func talk_comments_count(doc *goquery.Document) string{

	talk_comments_count := doc.Find(".h11").Contents().Text()
	//fmt.Println(talk_comments_count)
	d := strings.Split(talk_comments_count, " ")
	//fmt.Println(d[0])
	return strings.TrimLeft(d[0],"\n")
}

// this should return an array of strings => ["langs"]
func langs(doc *goquery.Document) []string {

	var langsList []string

	langs := doc.Find(".talk-transcript__language").Contents().Text()

	//fmt.Println(langs)
	langsSeparated := strings.Split(langs, "\n")

	for i := 1; i < len(langsSeparated)-1; i++ {
		//fmt.Println(i, ":", langsSeparated[i])
		langsList = append(langsList, langsSeparated[i])
	}

	return langsList
}

func main() {

	var doc []*goquery.Document

	//url := "https://www.ted.com/talks/ari_wallach_3_ways_to_plan_for_the_very_long_term/transcript?language=en"
	url := "https://www.ted.com/talks/ari_wallach_3_ways_to_plan_for_the_very_long_term"


	page, _ := goquery.NewDocument(url)
	doc = append(doc, page)


	//fmt.Println(doc[0])

	//fmt.Println(langs(doc)[1])

/*
	
			var availableLangs []string

			for _, code := range langs(doc[0]) {

				if langCodes[code] == "" {
					//fmt.Println(code, " : Not available")
				} else {
					//fmt.Println(langCodes[code])
					availableLangs = append(availableLangs, langCodes[code])
				}
			}

		fmt.Println(availableLangs)
	
for _,x := range availableLangs{
	transcriptURL := mainURL + "/transcript?language=" + x

	// asynchronously send the requests and fill in the correct place in the struct
	fmt.Println(transcriptURL)
}

*/

kenJsonObj := gabs.New()

// Video page
kenJsonObj.Set(url,"VideoPage", "TalkLink")

//kenJsonObj.Set(availableSubtitles(page),"VideoPage", "AvailableSubtitlesCount")

kenJsonObj.Set(speaker(page),"VideoPage", "Speaker")
kenJsonObj.Set(duration(page),"VideoPage", "Duration")
kenJsonObj.Set(time_filmed(page),"VideoPage", "TimeFilmed")
kenJsonObj.Set(talk_views_count(page),"VideoPage", "TalkViewsCount")
kenJsonObj.Set(talk_topics_list(page),"VideoPage", "TalkTopicsList")
kenJsonObj.Set(talk_comments_count(page),"VideoPage", "TalkCommentsCount")


// Transcript page
//kenJsonObj.Set(title(page),"TranscriptPage","TalkTitle")

fmt.Println(kenJsonObj.StringIndent(" ", "  "))



//fmt.Println(title(page))
ioutil.WriteFile("./experiments/goKEN.json", []byte(kenJsonObj.StringIndent(" ", "  ")), 0777)

}
