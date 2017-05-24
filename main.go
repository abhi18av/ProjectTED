package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

type TedTalk struct {
	TalkVideoPage      VideoPage      `json:"TalkVideoPage"`
	TalkTranscriptPage TranscriptPage `json:"TalkTranscriptPage"`
}

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

type talkTranscript struct {
	LocalTalkTitle              string   `json:"LocalTalkTitle"`
	Paragraphs                  []string `json:"Paragraphs"`
	TimeStamps                  []string `json:"TimeStamps"`
	TalkTranscriptAndTimeStamps []string `json:"TalkTranscriptAndTimeStamps"`
}

type TranscriptPage struct {
	AvailableTranscripts []string                  `json:"AvailableTranscripts"`
	DatePosted           string                    `json:"DatePosted"`
	Rated                string                    `json:"Rated"`
	TalkTranscript       map[string]talkTranscript `json:"TalkTranscript"`
}

func main() {

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func checkInternet() {
	// Make a GET request
	rs, err := http.Get("https://google.com")
	// Process response
	if err != nil {
		color.Red("We're OFF-Line!")
		//panic("Not connected to the net") // More idiomatic way would be to print the error and die unless it's a serious error

		// Learn about exit status in Golang
		os.Exit(1)
	}

	defer rs.Body.Close()

}

func exitIfNoSubtitlesExist(numOfSubtitles int64) {
	if numOfSubtitles < 1 {
		color.Red("No subtitles available yet")
		os.Exit(1)
	}
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

func transcriptFetchCommonInfo(url string) TranscriptPage {
	transcriptPage, _ := goquery.NewDocument(url)

	transcriptPageInstance := TranscriptPage{

		AvailableTranscripts: transcriptAvailableTranscripts(transcriptPage),
		DatePosted:           transcriptDatePosted(transcriptPage),
		Rated:                transcriptRated(transcriptPage),
	}
	return transcriptPageInstance
}

func transcriptFetchUncommonInfo(url string) (talkTranscript, string) {

	//fmt.Println(url)
	transcriptPage, _ := goquery.NewDocument(url)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	transcript := talkTranscript{

		LocalTalkTitle:              transcriptLocalTalkTitle(transcriptPage),
		Paragraphs:                  transcriptTalkTranscript(transcriptPage),
		TimeStamps:                  transcriptTimeStamps(transcriptPage),
		TalkTranscriptAndTimeStamps: transcriptTalkTranscriptAndTimeStamps(transcriptPage),
	}

	langName := strings.Split(url, "=")[1]
	return transcript, langName
}
