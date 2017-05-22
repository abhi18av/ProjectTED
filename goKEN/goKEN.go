package main

import (
	"net/http"
	"os"
	"strings"

	"sync"

	"encoding/json"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

type VideoPage struct {
	TalkURL                 string   `json:"VideoURL"`
	AvailableSubtitlesCount string   `json:"AvailableSubtitlesCount"`
	Speaker                 string   `json:"Speaker"`
	Duration                string   `json:"Duration"`
	TimeFilmed              string   `json:"TimeFilmed"`
	TalkViewsCount          string   `json:"TalkViewsCount"`
	TalkTopicsList          []string `json:"TalkTopicsList"`
	TalkCommentsCount       string   `json:"TalkCommentsCount"`
}

func main() {

	checkInternet()

	videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	var wg sync.WaitGroup
	var videoPageInfo VideoPage
	wg.Add(1)

	go func(videoURL string) {
		defer wg.Done()
		videoPageInfo = videoFetchInfo(videoURL)

	}(videoURL)

	wg.Wait()

	writeJSON(videoPageInfo)
}

func writeJSON(videoPageInfo VideoPage) {

	temp1, _ := json.Marshal(videoPageInfo)
	//fmt.Println(string(temp1))
	htmlSplit := strings.Split(videoPageInfo.TalkURL, "/")
	talkName := htmlSplit[len(htmlSplit)-1]

	fileName := "./" + talkName + ".json"

	f, err := os.Create(fileName)

	f.Write(temp1)
	check(err)
	defer f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkInternet() {
	// Make a GET request
	rs, err := http.Get("https://google.com")
	// Process response
	if err != nil {
		color.Red("OFF-Line")
		//panic("Not connected to the net") // More idiomatic way would be to print the error and die unless it's a serious error

		// Learn about exit status in Golang
		os.Exit(1)
	}

	defer rs.Body.Close()

}

func videoFetchInfo(url string) VideoPage {

	videoPage, _ := goquery.NewDocument(url)

	videoPageInstance := VideoPage{
		TalkURL:                 videoTalkURL(url),
		AvailableSubtitlesCount: videoAvailableSubtitlesCount(videoPage),
		Speaker:                 videoSpeaker(videoPage),
		Duration:                videoDuration(videoPage),
		TimeFilmed:              videoTimeFilmed(videoPage),
		TalkViewsCount:          videoTalkViewsCount(videoPage),
		TalkTopicsList:          videoTalkTopicsList(videoPage),
		TalkCommentsCount:       videoTalkCommentsCount(videoPage),
	}

	return videoPageInstance
}

// VIDEO functions

func videoAvailableSubtitlesCount(doc *goquery.Document) string {

	subtitles := doc.Find(".player-hero__meta__link").Contents().Text()
	//fmt.Println(subtitles)

	//for _, x := range strings.Split(subtitles, "\n") {
	//fmt.Println(x)
	//println("~~~~~~")
	//}

	y := strings.Split(subtitles, "\n")
	z := strings.Split(y[3], " ")[0]
	// In case I need an INT
	//numOfSubtitles, _ := strconv.ParseInt(z, 10, 32)
	numOfSubtitles := z
	return numOfSubtitles
}

func videoSpeaker(doc *goquery.Document) string {
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

func videoDuration(doc *goquery.Document) string {

	duration := doc.Find(".player-hero__meta").Contents().Text()
	//fmt.Println(duration)

	//for _, x := range strings.Split(duration, "\n") {
	//	fmt.Println(x)
	//	println("~~~~~~")
	//}

	x := strings.Split(duration, "\n")
	//fmt.Println(x[6])
	return x[6]

}

// TimeFilmed : Time at which the talk was filmed
func videoTimeFilmed(doc *goquery.Document) string {

	time_filmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(time_filmed)

	y := strings.Split(time_filmed, "\n")
	//fmt.Println(y[11])
	return y[11]
}

func videoTalkViewsCount(doc *goquery.Document) string {

	talk_views_count := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talk_views_count)

	a := strings.Split(talk_views_count, "\n")
	b := strings.TrimSpace(a[2])
	//fmt.Println(b)
	return b

}

func videoTalkTopicsList(doc *goquery.Document) []string {

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

func videoTalkCommentsCount(doc *goquery.Document) string {

	talk_comments_count := doc.Find(".h11").Contents().Text()
	//fmt.Println(talk_comments_count)
	d := strings.Split(talk_comments_count, " ")
	//fmt.Println(d[0])
	return strings.TrimLeft(d[0], "\n")
}

func videoTalkURL(videoURL string) string {
	return videoURL
}
