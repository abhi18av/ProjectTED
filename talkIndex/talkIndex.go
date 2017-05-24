package main

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"encoding/json"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {

	startingIndexURL := "https://www.ted.com/talks?page=1"
	page, _ := goquery.NewDocument(startingIndexURL)

	lastIndex := lastIndex(page)
	numLastIndex, _ := strconv.ParseInt(lastIndex, 10, 64)
	intLastIndex := int(numLastIndex)
	//fmt.Println(numLastIndex)

	//linksInAPage := collectTalkLinks(page)
	//fmt.Println(linksInAPage)

	mainTedURL := "https://www.ted.com"

	var allTalksLinks [][]string
	var wg sync.WaitGroup
	wg.Add(intLastIndex)
	for i := 1; i <= intLastIndex; i++ {
		URL := mainTedURL + "/talks?page=" + strconv.FormatInt(int64(i), 10)
		color.Blue(URL)

		//fmt.Println("Page : ", index)

		go func(URL string) {
			defer wg.Done()
			aPage, _ := goquery.NewDocument(URL)
			links := collectTalkLinks(aPage)

			allTalksLinks = append(allTalksLinks, links)

			color.Green(URL)
		}(URL)

	}

	wg.Wait()

	outPut := flattenStringMatrix(allTalksLinks)
	writeJSON(outPut)

	//fmt.Println(len(allTalksLinks))
	//x, _ := json.Marshal(collectTalkLinks)
	//fmt.Println(x)

}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func collectTalkLinks(doc *goquery.Document) []string {
	var urls []string

	doc.Find(".h9").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		//fmt.Println(i+1, link)
		urls = append(urls, link)
	})

	//fmt.Println(urls)
	return urls
}

func lastIndex(doc *goquery.Document) string {
	x := doc.Find(".pagination__flipper").PrevUntil(".pagination__gap").Contents().Text()
	y := strings.Split(x, "|")
	//fmt.Println(y[1])
	return y[1]
}

func writeJSON(aStruct [][]string) {

	temp1, _ := json.Marshal(aStruct)

	f, err := os.Create("./out.txt")
	checkErr(err)

	f.Write(temp1)
	defer f.Close()
}

//twoDArr := [][]string{[]string{"a", "b", "c"}, []string{"x", "y", "z"}}
//flattenStringMatrix(twoDArr)

func flattenStringMatrix(aMatrix [][]string) []string {

	var arr []string

	for _, anArr := range aMatrix {

		for i := 0; i < len(anArr); i++ {

			arr = append(arr, anArr[i])

		}

	}

	//fmt.Println(arr)
}
