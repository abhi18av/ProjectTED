package main

import (
	"github.com/PuerkitoBio/goquery"
)

func main() {
	videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	//transcriptsURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	videoURL, _ := goquery.NewDocument(videoURL)

	videoPage.

}


// VideoPage FUNCTIONS


var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
	"German":              "de",
	"Russian":             "ru",
}
func AvailableSubtitlesCount(doc *goquery.Document) string {

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

// transcriptPage 


func DatePosted(doc *goquery.Document) string {
	posted := doc.Find(".meta__item").Contents().Text()
	p := strings.Split(posted, "\n")
	//fmt.Println(p[3])
	return (p[3])

}

func Rated(doc *goquery.Document) string {

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

func LocalTalkTitle(doc *goquery.Document) string {
	title := doc.Find(".m5").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}

func TimeStamps(doc *goquery.Document) []string {
	times := doc.Find(".talk-transcript__para__time").Contents().Text()
	var timestamps []string

	for _, time := range strings.Split(times, " ") {
		if time == "" {

		} else {

			fmt.Println(time)
			timestamps = append(timestamps, strings.TrimRight(time, "\n"))

		}
	}
	return timestamps
}

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

// this should return an array of strings => ["langs"]
func AvailableTranscripts(doc *goquery.Document) []string {

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
