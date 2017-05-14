package main

import (
	"github.com/Jeffail/gabs"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	video_url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"

	transcripts_url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"

	video_page, _ := goquery.NewDocument(video_url)

	kenJsonObj := gabs.New()

	// Video page
	kenJsonObj.Set(video_url, "VideoPage", "TalkLink")
	kenJsonObj.Set(availableSubtitlesCount(video_page), "VideoPage", "AvailableSubtitlesCount")
	kenJsonObj.Set(speaker(video_page), "VideoPage", "Speaker")
	kenJsonObj.Set(duration(video_page), "VideoPage", "Duration")
	kenJsonObj.Set(time_filmed(video_page), "VideoPage", "TimeFilmed")
	kenJsonObj.Set(talk_views_count(video_page), "VideoPage", "TalkViewsCount")
	kenJsonObj.Set(talk_topics_list(video_page), "VideoPage", "TalkTopicsList")
	kenJsonObj.Set(talk_comments_count(video_page), "VideoPage", "TalkCommentsCount")

}
