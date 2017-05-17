package main

import (
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

type VideoPage struct {
	AvailableSubtitlesCount string   `json:"AvailableSubtitlesCount"`
	Speaker                 string   `json:"Speaker"`
	Duration                string   `json:"Duration"`
	TimeFilmed              string   `json:"TimeFilmed"`
	TalkViewsCount          string   `json:"TalkViewsCount"`
	TalkTopicsList          []string `json:"TalkTopicsList"`
	TalkCommentsCount       string   `json:"TalkCommentsCount"`
}

type TranscriptPage struct {
	AvailableTranscripts []string `json:"AvailableTranscripts"`
	DatePosted           string   `json:"DatePosted"`
	LocalTitle           string   `json:"LocalTitle"`
	Rated                string   `json:"Rated"`
	TalkTranscript       talkTranscript
	TimeStamps           []string `json:"TimeStamps"`
}
type talkTranscript struct {
	LocalTalkTitle string   `json:"LocalTalkTitle"`
	Paragraphs     []string `json:"Paragraphs"`
}

func main() {

	/*
		// VIDEO functions
		videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"
		videoPage, _ := goquery.NewDocument(videoURL)

		//fmt.Println(videoTalkTitle(videoPage))


			videoPageInstance := VideoPage{
				AvailableSubtitlesCount: videoAvailableSubtitlesCount(videoPage),

				Speaker:           videoSpeaker(videoPage),
				Duration:          videoDuration(videoPage),
				TimeFilmed:        videoTimeFilmed(videoPage),
				TalkViewsCount:    videoTalkViewsCount(videoPage),
				TalkTopicsList:    videoTalkTopicsList(videoPage),
				TalkCommentsCount: videoTalkCommentsCount(videoPage),
			}

			body, err := json.Marshal(videoPageInstance)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))

	*/

	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

	// TRANSCRIPT functions
	//transcriptURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=de"
	//transcriptPage, _ := goquery.NewDocument(transcriptURL)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	var wg sync.WaitGroup

	var urls []string

	transcriptBaseURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"
	langBaseURL := "/transcript?language="

	for _, value := range langCodes {
		newURL := transcriptBaseURL + langBaseURL + value
		//fmt.Println(x)
		urls = append(urls, newURL)
	}
	//fmt.Println(len(urls))

	//fmt.Println(urls)

	wg.Add(len(urls))

	for _, url := range urls {

		go func(url string) {

			//fmt.Println(url)
			transcriptPage, _ := goquery.NewDocument(url)
			//fmt.Println(transcriptLocalTalkTitle(transcriptPage))
			transcriptPageInstance := TranscriptPage{

				LocalTitle: transcriptLocalTalkTitle(transcriptPage),
			}
			fmt.Println(transcriptPageInstance)
			defer wg.Done()
		}(url)
	}

	wg.Wait()

} // end of main()

// VideoPage FUNCTIONS

// OUTPUT
// 60
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

// OUTPUT
// Ken Robinson
func videoSpeaker(doc *goquery.Document) string {
	speaker := doc.Find(".talk-speaker__name").Contents().Text()
	//fmt.Println(speaker)
	speaker = strings.Trim(speaker, "\n")
	return speaker
}

// This is now taken from the transcripts page
func videoTalkTitle(doc *goquery.Document) string {
	title := doc.Find(".player-hero__title__content").Contents().Text()
	//fmt.Println(title)
	return title
}

// OUTPUT
// 19:24
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

// OUTPUT
// Feb 2006
// TimeFilmed : Time at which the talk was filmed
func videoTimeFilmed(doc *goquery.Document) string {

	timeFilmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(timeFilmed)

	y := strings.Split(timeFilmed, "\n")
	//fmt.Println(y[11])
	return y[11]
}

// OUTPUT
// 45,122,067
func videoTalkViewsCount(doc *goquery.Document) string {

	talkViewsCount := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talkViewsCount)

	a := strings.Split(talkViewsCount, "\n")
	b := strings.TrimSpace(a[2])
	//fmt.Println(b)
	return b

}

// OUTPUT
// [Children Creativity Culture Dance Education Parenting Teaching]
func videoTalkTopicsList(doc *goquery.Document) []string {

	talkTopics := doc.Find(".talk-topics__list").Contents().Text()

	c := strings.Split(talkTopics, "\n")
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

// OUTPUT
// 4526
func videoTalkCommentsCount(doc *goquery.Document) string {

	talkCommentsCount := doc.Find(".h11").Contents().Text()
	//fmt.Println(talkCommentsCount)
	d := strings.Split(talkCommentsCount, " ")
	//fmt.Println(d[0])
	return strings.TrimLeft(d[0], "\n")
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
