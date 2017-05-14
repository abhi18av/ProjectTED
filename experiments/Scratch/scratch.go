package main

import (
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/PuerkitoBio/goquery"
)

type TranscriptPage struct {
	LocalTalkTitle string
	TalkTranscript []string
}

func LocalTalkTitle(doc *goquery.Document) string {
	title := doc.Find(".m5").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}

func TalkTranscript(doc *goquery.Document) []string {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	var para []string
	for _, text := range strings.Split(texts, "  ") {

		//fmt.Println(text)
		para = append(para, text)
	}
	return para
}

func main() {

	transcriptURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	transcriptPage, _ := goquery.NewDocument(videoURL)

	kenJsonObj := gabs.New()

	// Video page
	kenJsonObj.Set(videoURL, "transcriptPage", "TalkLink")
	kenJsonObj.Set(availableSubtitlesCount(transcriptPage), "transcriptPage", "AvailableSubtitlesCount")
	kenJsonObj.Set(talk_topics_list(), "transcriptPage", "TalkTopicsList")
	kenJsonObj.Set(talk_comments_count(), "transcriptPage", "TalkCommentsCount")

}
