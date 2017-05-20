package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
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
	LocalTalkTitle string   `json:"LocalTalkTitle"`
	Paragraphs     []string `json:"Paragraphs"`
	TimeStamps     []string `json:"TimeStamps"`
}

type TranscriptPage struct {
	AvailableTranscripts []string         `json:"AvailableTranscripts"`
	DatePosted           string           `json:"DatePosted"`
	Rated                string           `json:"Rated"`
	TalkTranscript       []talkTranscript `json:"TalkTranscript"`
}

func main() {
	// TRANSCRIPT functions
	//transcriptURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=de"
	//transcriptPage, _ := goquery.NewDocument(transcriptURL)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	urls := genTranscriptURLs(langCodes, videoURL)

	//fmt.Println(urls)

	var wg sync.WaitGroup
	wg.Add(len(urls))
	var transcript talkTranscript
	var transcriptPageInstance TranscriptPage

	for _, url := range urls {

		go func(url string, transcript talkTranscript, transcriptPageInstance TranscriptPage) {

			//fmt.Println(url)
			transcriptPage, _ := goquery.NewDocument(url)
			//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

			transcript = talkTranscript{

				LocalTalkTitle: transcriptLocalTalkTitle(transcriptPage),
				Paragraphs:     transcriptTalkTranscript(transcriptPage),
				TimeStamps:     transcriptTimeStamps(transcriptPage),
			}

			// Using append here to add to the array-field
			transcriptPageInstance.TalkTranscript = append(transcriptPageInstance.TalkTranscript, transcript)

			transcriptPageInstance = TranscriptPage{

				AvailableTranscripts: transcriptAvailableTranscripts(transcriptPage),
				DatePosted:           transcriptDatePosted(transcriptPage),
				Rated:                transcriptRated(transcriptPage),
				//TalkTranscript:       transcript,
			}

			//fmt.Println(transcriptPageInstance)
			//fmt.Println(transcript)

			body, err := json.Marshal(transcriptPageInstance)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))

			defer wg.Done()
		}(url, transcript, transcriptPageInstance)
	}

	wg.Wait()

} // end of main()

func fetch(url string, ch chan talkTranscript) talkTranscript {

	//fmt.Println(url)
	transcriptPage, _ := goquery.NewDocument(url)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	transcript := talkTranscript{

		LocalTalkTitle: transcriptLocalTalkTitle(transcriptPage),
		Paragraphs:     transcriptTalkTranscript(transcriptPage),
		TimeStamps:     transcriptTimeStamps(transcriptPage),
	}
	//fmt.Println(transcript)
	return transcript
}

// transcriptPage

// OUTPUT
// Jun 2006
func transcriptDatePosted(doc *goquery.Document) string {
	posted := doc.Find(".meta__item").Contents().Text()
	p := strings.Split(posted, "\n")
	//fmt.Println(p[3])
	return (p[3])

}

// OUTPUT
// Inspiring, Funny
func transcriptRated(doc *goquery.Document) string {

	rated := doc.Find(".meta__row").Contents().Text()

	r := strings.Split(rated, "\n")
	//fmt.Println(r[3])
	return r[3]
	/*
	   rx := strings.Split(r[3], ",")

	   	for _, x := range rx{
	   		append(ls,x)
	   	}
	*/

	//println(len(rx))
	//println(r[0])
	//println(r[1])
	//return(p[3])
}
