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
}

type TranscriptPage struct {
	AvailableTranscripts []string         `json:"AvailableTranscripts"`
	DatePosted           string           `json:"DatePosted"`
	Rated                string           `json:"Rated"`
	TalkTranscript       []talkTranscript `json:"TalkTranscript"`
	TimeStamps           []string         `json:"TimeStamps"`
}

func main() {
	// TRANSCRIPT functions
	//transcriptURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=de"
	//transcriptPage, _ := goquery.NewDocument(transcriptURL)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	urls := genTranscriptURLs(langCodes, videoURL)

	fmt.Println(urls)

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {

		go func(url string) {

			//fmt.Println(url)
			transcriptPage, _ := goquery.NewDocument(url)
			//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

			transcript := talkTranscript{

				LocalTalkTitle: transcriptLocalTalkTitle(transcriptPage),
				Paragraphs:     transcriptTalkTranscript(transcriptPage),
			}

			transcriptPageInstance := TranscriptPage{

				AvailableTranscripts: transcriptAvailableTranscripts(transcriptPage),
				DatePosted:           transcriptDatePosted(transcriptPage),
				Rated:                transcriptRated(transcriptPage),
				//TalkTranscript:       transcript,
				TimeStamps: transcriptTimeStamps(transcriptPage),
			}

			// Using append here to add to the array-field
			transcriptPageInstance.TalkTranscript = append(transcriptPageInstance.TalkTranscript, transcript)

			//fmt.Println(transcriptPageInstance)
			//fmt.Println(transcript)

			body, err := json.Marshal(transcriptPageInstance)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))

			defer wg.Done()
		}(url)
	}

	wg.Wait()

} // end of main()

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

// OUTPUT
// Do schools kill creativity?
func transcriptLocalTalkTitle(doc *goquery.Document) string {
	title := doc.Find(".m5").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}

/*

Need to learn from this function



func TalkTranscript(doc *goquery.Document) []string {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	var para []string
	for _, text := range strings.Split(texts, "  ") {

		//fmt.Println(text)
		para = append(para, text)
	}

	var lines []string
	for _, para := range strings.Split(texts, "\n\n") {

		//fmt.Println(text)
		lines = append(lines, para)
	}

	return para
	//return lines
}


*/

func transcriptTimeStamps(doc *goquery.Document) []string {
	times := doc.Find(".talk-transcript__para__time").Contents().Text()
	var timestamps []string

	for _, time := range strings.Split(times, " ") {
		if time == "" {

		} else {

			//fmt.Println(time)
			timestamps = append(timestamps, strings.TrimRight(time, " "))

		}
	}
	return timestamps
}

func transcriptTalkTranscript(doc *goquery.Document) []string {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	var para []string
	for _, text := range strings.Split(texts, "  ") {

		//fmt.Println(text)
		para = append(para, text)
	}

	var lines []string
	for _, para := range strings.Split(texts, "\n\n") {

		//fmt.Println(text)
		lines = append(lines, para)
	}

	return para
	//return lines
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
