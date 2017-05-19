package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
type TranscriptPage struct {
	LocalTalkTitle string
	TalkTranscript []string
}
*/

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

	for i := 0; i < len(times1)-1; i++ {

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
func main() {

	transcriptURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	transcriptPage, _ := goquery.NewDocument(transcriptURL)

	transcriptTimeStamps(transcriptPage)
}
