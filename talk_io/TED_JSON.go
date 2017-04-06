package main

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	"os"
)

type TED_JSON struct {
	TalkTitle  string   `json:"talk_title"`
	TalkTopics []string `json:"talk_topics"`
	TalkLink   string   `json:"talk_link"`
	Time       []string `json:"time"`
	UpdateDate string   `json:"update_date"`
	PostedDate string   `json:"posted_date"`
	Transcript struct {
		Te   []string `json:"te"`
		My   []string `json:"my"`
		Ko   []string `json:"ko"`
		Pt   []string `json:"pt"`
		Sk   []string `json:"sk"`
		Mk   []string `json:"mk"`
		Ro   []string `json:"ro"`
		Fa   []string `json:"fa"`
		Sr   []string `json:"sr"`
		He   []string `json:"he"`
		Sv   []string `json:"sv"`
		Es   []string `json:"es"`
		Fi   []string `json:"fi"`
		ZhTw []string `json:"zh-tw"`
		ZhCn []string `json:"zh-cn"`
		De   []string `json:"de"`
		PtBr []string `json:"pt-br"`
		Uk   []string `json:"uk"`
		Cs   []string `json:"cs"`
		Tr   []string `json:"tr"`
		Ru   []string `json:"ru"`
		It   []string `json:"it"`
		Eo   []string `json:"eo"`
		Bg   []string `json:"bg"`
		Hi   []string `json:"hi"`
		El   []string `json:"el"`
		Ar   []string `json:"ar"`
		Fr   []string `json:"fr"`
		Hr   []string `json:"hr"`
		Be   []string `json:"be"`
		Hu   []string `json:"hu"`
		En   []string `json:"en"`
		Vi   []string `json:"vi"`
		Nl   []string `json:"nl"`
		Th   []string `json:"th"`
		Ja   []string `json:"ja"`
		Mr   []string `json:"mr"`
	} `json:"transcript"`
}

func (p TED_JSON) toString() string {
	return toJson(p)
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
func main() {
	raw, _ := ioutil.ReadFile("./al-A_simple_way_to_break_a_bad_habit.json")
	/*
		var TED_JSONs []TED_JSON
		json.Unmarshal(raw, &TED_JSONs)

		for _, p := range TED_JSONs {
			fmt.Println(p.toString())
		}

		fmt.Println(toJson(TED_JSONs))
	*/
	jsonParsed, _ := gabs.ParseJSON(raw)

	x, _ := jsonParsed.Path("transcript.zh-cn").Children()

	for _, child := range x {

		fmt.Println(child.Data())
	}

}
