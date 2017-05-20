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

// OUTPUT
// Do schools kill creativity?
func transcriptLocalTalkTitle(doc *goquery.Document) string {
	title := doc.Find(".m5").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}

// OUTPUT
// ["0:11","0:15","0:16","0:23","0:29","0:56","1:11","1:15","1:18","1:21","1:35","1:37","1:38","1:41","2:23","2:56","3:09","3:11","3:16","3:18","3:20","3:22","3:25","3:27","3:30","3:56","4:07","4:12","4:14","4:20","4:21","4:25","4:27","5:09","5:21","6:05","6:21","6:31","6:33","6:54","7:01","7:03","7:10","7:11","7:15","7:21","7:22","7:24","7:28","7:29","7:34","7:57","7:58","8:10","8:18","8:21","8:27","9:10","9:13","9:22","9:48","9:51","10:21","10:27","10:30","10:36","10:45","10:48","10:54","10:56","11:00","11:02","11:18","11:41","12:06","12:23","12:56","13:33","13:56","13:58","14:18","14:25","14:26","14:28","14:43","14:50","15:46","15:48","15:50","15:53","16:26","16:50","17:32","17:39","18:13","18:33","19:04","19:05"]
func transcriptTimeStamps(doc *goquery.Document) []string {
	times := doc.Find(".talk-transcript__para__time").Contents().Text()
	var times1 []string

	for _, time := range strings.Split(times, " ") {
		if time == "" {

		} else {

			//fmt.Println(time)
			times1 = append(times1, strings.TrimRight(time, " "))

		}
	}

	//fmt.Println(times1)
	var times2 []string
	for _, time := range strings.Split(times1[len(times1)-1], "\n\n") {

		times2 = append(times2, time)
	}

	//fmt.Println(times2)
	var timestamps []string

	for i := 1; i < len(times1)-1; i++ {

		timestamps = append(timestamps, times1[i])

	}

	for i := 0; i < len(times2); i++ {

		timestamps = append(timestamps, times2[i])
	}

	//fmt.Println(timestamps)

	for i := 0; i < len(timestamps); i++ {

		timestamps[i] = strings.Trim(timestamps[i], "\n")
	}

	//x, _ := json.Marshal(timestamps)
	//fmt.Println(string(x))

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

// WORK ON THIS
func transcriptTalkTranscriptAndTimeStamps(doc *goquery.Document) {

	title := doc.Find(".talk-transcript__para").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}