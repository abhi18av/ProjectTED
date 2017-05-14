package main

type Foo struct {
	Meta struct {
		DateProcessed struct {
			ItemsUpdated struct {
				TranscriptPage interface{} `json:"TranscriptPage"`
				VideoPage      interface{} `json:"VideoPage"`
			} `json:"ItemsUpdated"`
		} `json:"DateProcessed"`
	} `json:"Meta"`
	TranscriptPage struct {
		AvailableTranscripts interface{} `json:"AvailableTranscripts"`
		DatePosted           interface{} `json:"DatePosted"`
		LocalTime            interface{} `json:"LocalTime"`
		Rated                interface{} `json:"Rated"`
		TalkTranscript       struct {
			LangCode struct {
				LocalTalkTitle interface{} `json:"LocalTalkTitle"`
				Paragraphs     interface{} `json:"Paragraphs"`
			} `json:"LangCode"`
		} `json:"TalkTranscript"`
		TimeStamps interface{} `json:"TimeStamps"`
	} `json:"TranscriptPage"`
	VideoPage struct {
		AvailableSubtitlesCount interface{} `json:"AvailableSubtitlesCount"`
		Duration                interface{} `json:"Duration"`
		Speaker                 interface{} `json:"Speaker"`
		TalkCommentsCount       interface{} `json:"TalkCommentsCount"`
		TalkLink                interface{} `json:"TalkLink"`
		TalkTitle               interface{} `json:"TalkTitle"`
		TalkTopicsList          interface{} `json:"TalkTopicsList"`
		TalkViewsCount          interface{} `json:"TalkViewsCount"`
		TimeFilmed              interface{} `json:"TimeFilmed"`
	} `json:"VideoPage"`
}
