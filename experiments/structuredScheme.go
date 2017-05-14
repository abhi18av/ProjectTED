package main

type Meta struct {
	DateProcessed []string
}

type TranscriptPage struct {
	AvailableTranscripts []string
	DatePosted           string
	LocalTitle           string
	Rated                string
	TalkTranscript       struct {
		LocalTalkTitle string
		Paragraphs     []string
	}
	TimeStamps []string
}

type VideoPage struct {
	AvailableSubtitlesCount string
	Speaker                 string
	Duration                string
	TimeFilmed              string
	TalkViewsCount          string
	TalkTopicsList          []string
	TalkCommentsCount       []string
}
