package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/imdario/mergo"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "fr",
	"German":              "de",
	"Russian":             "ru",
}

func genTranscriptURLs(langCodes map[string]string, videoURL string) []string {

	langBaseURL := "/transcript?language="

	var urls []string

	for _, value := range langCodes {
		newURL := videoURL + langBaseURL + value
		//fmt.Println(x)
		urls = append(urls, newURL)
	}
	//fmt.Println(len(urls))

	return urls
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

func WiFi() {
	// Make a get request
	rs, err := http.Get("https://google.com")
	// Process response
	if err != nil {
		color.Red("WiFi OFF")
		//panic("Not connected to the net") // More idiomatic way would be to print the error and die unless it's a serious error

		// Learn about exit status in Golang
		os.Exit(1)
	}

	defer rs.Body.Close()

}

func main() {

	WiFi()

	videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"
	urls := genTranscriptURLs(langCodes, videoURL)

	//fmt.Println(urls)

	var wg sync.WaitGroup
	wg.Add(len(urls) + 1)

	// @@@@@@@@@@
	// Page Common
	transcriptEnURL := videoURL + "/transcript?language=en"

	var transcriptPageCommon TranscriptPage

	go func(url string) {
		defer wg.Done()
		transcriptPageCommon = fetchCommon(url)
	}(transcriptEnURL)

	// @@@@@@@@@@
	// Page UnCommon

	var transcriptS []talkTranscript

	langSpecificMap := make(map[string]talkTranscript)

	for _, url := range urls {

		go func(url string) {
			defer wg.Done()
			x, langName := fetchUncommon(url)
			//color.Blue(langName)

			langSpecificMap[langName] = x
			transcriptS = append(transcriptS, x)
			//transcriptS.TalkTranscript = langSpecificMap
		}(url)

	}

	// @@@@@@@@@@@@
	wg.Wait()

	var transcriptPageUnCommon TranscriptPage
	//fmt.Println(langSpecificMap)
	transcriptPageUnCommon.TalkTranscript = langSpecificMap
	transcriptPageComplete := transcriptPageCommon

	mergo.Merge(&transcriptPageComplete, transcriptPageUnCommon)
	z, _ := json.Marshal(transcriptPageComplete)
	htmlSplit := strings.Split(videoURL, "/")
	talkName := htmlSplit[len(htmlSplit)-1]
	fmt.Println(talkName)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetchCommon(url string) TranscriptPage {
	transcriptPage, _ := goquery.NewDocument(url)

	// Using append here to add to the array-field
	//transcriptPageInstance.TalkTranscript = append(transcriptPageInstance.TalkTranscript, transcript)

	transcriptPageInstance := TranscriptPage{

		AvailableTranscripts: transcriptAvailableTranscripts(transcriptPage),
		DatePosted:           transcriptDatePosted(transcriptPage),
		Rated:                transcriptRated(transcriptPage),
		//TalkTranscript:       transcriptS,
	}
	return transcriptPageInstance
}

func fetchUncommon(url string) (talkTranscript, string) {

	//fmt.Println(url)
	transcriptPage, _ := goquery.NewDocument(url)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	transcript := talkTranscript{

		LocalTalkTitle:              transcriptLocalTalkTitle(transcriptPage),
		Paragraphs:                  transcriptTalkTranscript(transcriptPage),
		TimeStamps:                  transcriptTimeStamps(transcriptPage),
		TalkTranscriptAndTimeStamps: transcriptTalkTranscriptAndTimeStamps(transcriptPage),
	}
	//fmt.Println(transcript)

	langName := strings.Split(url, "=")[1]
	//color.Blue(langName)
	return transcript, langName
}

// transcriptPage

// OUTPUT
// Do schools kill creativity?
func transcriptLocalTalkTitle(doc *goquery.Document) string {
	title := doc.Find(".m5").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}

// OUTPUT
// [Afrikaans Albanian Arabic Armenian Azerbaijani Basque Belarusian Bengali Bulgarian Catalan Chinese, Simplified Chinese, Traditional Croatian Czech Danish Dutch English Esperanto Estonian Filipino Finnish French French (Canada) Galician Georgian German Greek Hebrew Hungarian Indonesian Ingush Italian Japanese Korean Lao Latvian Lithuanian Macedonian Marathi Mongolian Nepali Norwegian Bokmal Persian Polish Portuguese Portuguese, Brazilian Romanian Russian Serbian Slovak Slovenian Spanish Swedish Thai Turkish Ukrainian Urdu Uzbek Vietnamese]
// This should return an array of strings => ["langs"]
func transcriptAvailableTranscripts(doc *goquery.Document) []string {

	var langsList []string

	langs := doc.Find(".talk-transcript__language").Contents().Text()

	//	fmt.Println(langs)
	langsSeparated := strings.Split(langs, "\n")

	for i := 1; i < len(langsSeparated)-1; i++ {
		//fmt.Println(i, ":", langsSeparated[i])
		langsList = append(langsList, langsSeparated[i])
	}

	return langsList
}
