package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var langCodes = map[string]string{
	"Afrikaans":             "af",
	"Albanian":              "sq",
	"Algerian Arabic":       "arq",
	"Amharic":               "am",
	"Arabic":                "ar",
	"Armenian":              "hy",
	"Assamese":              "as",
	"Asturian":              "ast",
	"Azerbaijani":           "az",
	"Basque":                "eu",
	"Belarusian":            "be",
	"Bengali":               "bn",
	"Bislama":               "bi",
	"Bosnian":               "bs",
	"Bulgarian":             "bg",
	"Burmese":               "my",
	"Catalan":               "ca",
	"Cebuano":               "ceb",
	"Chinese, Simplified":   "zh-cn",
	"Chinese, Traditional":  "zh-tw",
	"Chinese, Yue":          "zh",
	"Creole, Haitian":       "ht",
	"Croatian":              "hr",
	"Czech":                 "cs",
	"Danish":                "da",
	"Dutch":                 "nl",
	"Dzongkha":              "dz",
	"English":               "en",
	"Esperanto":             "eo",
	"Estonian":              "et",
	"Filipino":              "fil",
	"Finnish":               "fi",
	"French":                "fr",
	"French (Canada)":       "fr-ca",
	"Galician":              "gl",
	"Georgian":              "ka",
	"German":                "de",
	"Greek":                 "el",
	"Gujarati":              "gu",
	"Hakha Chin":            "cnh",
	"Hausa":                 "ha",
	"Hebrew":                "he",
	"Hindi":                 "hi",
	"Hungarian":             "hu",
	"Hupa":                  "hup",
	"Icelandic":             "is",
	"Igbo":                  "ig",
	"Indonesian":            "id",
	"Ingush":                "inh",
	"Irish":                 "ga",
	"Italian":               "it",
	"Japanese":              "ja",
	"Kannada":               "kn",
	"Kazakh":                "kk",
	"Khmer":                 "km",
	"Klingon":               "tlh",
	"Korean":                "ko",
	"Kurdish":               "ku",
	"Kyrgyz":                "ky",
	"Lao":                   "lo",
	"Latgalian":             "ltg",
	"Latin":                 "la",
	"Latvian":               "lv",
	"Lithuanian":            "lt",
	"Luxembourgish":         "lb",
	"Macedo":                "rup",
	"Macedonian":            "mk",
	"Malagasy":              "mg",
	"Malay":                 "ms",
	"Malayalam":             "ml",
	"Maltese":               "mt",
	"Marathi":               "mr",
	"Mauritian Creole":      "mfe",
	"Mongolian":             "mn",
	"Montenegrin":           "srp",
	"Nepali":                "ne",
	"Norwegian Bokmal":      "nb",
	"Norwegian Nynorsk":     "nn",
	"Occitan":               "oc",
	"Pashto":                "ps",
	"Persian":               "fa",
	"Polish":                "pl",
	"Portuguese":            "pt",
	"Portuguese, Brazilian": "pt-br",
	"Punjabi":               "pa",
	"Romanian":              "ro",
	"Russian":               "ru",
	"Rusyn":                 "ry",
	"Serbian":               "sr",
	"Serbo-Croatian":        "sh",
	"Silesian":              "szl",
	"Sinhala":               "si",
	"Slovak":                "sk",
	"Slovenian":             "sl",
	"Somali":                "so",
	"Spanish":               "es",
	"Swahili":               "sw",
	"Swedish":               "sv",
	"Swedish Chef":          "art-x-bork",
	"Tagalog":               "tl",
	"Tajik":                 "tg",
	"Tamil":                 "ta",
	"Tatar":                 "tt",
	"Telugu":                "te",
	"Thai":                  "th",
	"Tibetan":               "bo",
	"Turkish":               "tr",
	"Turkmen":               "tk",
	"Ukrainian":             "uk",
	"Urdu":                  "ur",
	"Uyghur":                "ug",
	"Uzbek":                 "uz",
	"Vietnamese":            "vi",
}

func Posted(doc *goquery.Document) string {
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

func LocalTitle(doc *goquery.Document) string {
	title := doc.Find(".m5").Contents().Text()
	//fmt.Println(strings.Split(title, "\n")[2])
	return strings.Split(title, "\n")[2]
}

func Times(doc *goquery.Document) []string {
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

func TalkTexts(doc *goquery.Document) []string {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	var para []string
	for _, text := range strings.Split(texts, "  ") {

		//fmt.Println(text)
		para = append(para, text)
	}
	return para
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

func main() {

	//url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en"
	url := "https://www.ted.com/talks/jill_bolte_taylor_s_powerful_stroke_of_insight/transcript?language=en"
	//url := "https://www.ted.com/talks/tony_robbins_asks_why_we_do_what_we_do/transcript?language=en"
	//url := "https://www.ted.com/talks/shawn_achor_the_happy_secret_to_better_work/transcript?language=en"
	//url := "https://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action/transcript?language=en"

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	//	posted(doc)
	//	rated(doc)

	// Use this for the main title in English
	LocalTitle(doc)

	//println(len(langs(doc)))
	//	times(doc)
	//	texts(doc)

}
