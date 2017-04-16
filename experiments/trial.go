package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	//"os"
	"github.com/Jeffail/gabs"
)


type kenJSON struct {
	// This is the main title from English language
	TalkTitle string `json:"talk_title"`
	// The link to the main video
	TalkLink string `json:"talk_link"`
}

// From the main page or the English language transcript

func title(doc *goquery.Document) string {
	holder := doc.Find(".m5").Contents().Text()
	title := strings.Split(holder, "\n")[2]
	//fmt.Println(title)
	return title
}

func main() {
	// Need a check for internet connection 

	url := "https://www.ted.com/talks/ari_wallach_3_ways_to_plan_for_the_very_long_term/transcript?language=en"


	page, _ := goquery.NewDocument(url)


//info := kenJSON{ TalkLink: url, TalkTitle: title(page)}
//fmt.Println(info)

kenJsonObj := gabs.New()
kenJsonObj.Set(url, "TalkLink")
kenJsonObj.Set(title(page), "TalkTitle")
fmt.Println(kenJsonObj.String())

//fmt.Println(title(page))
ioutil.WriteFile("./output.json", []byte(kenJsonObj.String()), 0777)

}