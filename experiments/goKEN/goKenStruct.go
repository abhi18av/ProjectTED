package main


type goKEN{

TalkMeta Meta
TalkTranscriptPage TranscriptPage
TalkVideoPage VideoPage

}
type Meta struct {
	DateProcessed []string
	type ItemsUpdated struct{
   	TranscriptPage
      	VideoPage
	}
}

type TranscriptPage struct {
	AvailableTranscripts []string
	DatePosted           string
	LocalTitle           string
	Rated                string
	TalkTranscript       []string {
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
