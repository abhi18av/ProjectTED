func main() {

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
