package main

import (
	"encoding/json"
	"fmt"
	"os"
)




type TED_JSON struct {
	TalkTitle string `json:"talk_title"`
	TalkTopics []string `json:"talk_topics"`
	TalkLink string `json:"talk_link"`
	Time []string `json:"time"`
	UpdateDate string `json:"update_date"`
	PostedDate string `json:"posted_date"`
	Transcript struct {
		Te []string `json:"te"`
		My []string `json:"my"`
		Ko []string `json:"ko"`
		Pt []string `json:"pt"`
		Sk []string `json:"sk"`
		Mk []string `json:"mk"`
		Ro []string `json:"ro"`
		Fa []string `json:"fa"`
		Sr []string `json:"sr"`
		He []string `json:"he"`
		Sv []string `json:"sv"`
		Es []string `json:"es"`
		Fi []string `json:"fi"`
		ZhTw []string `json:"zh-tw"`
		ZhCn []string `json:"zh-cn"`
		De []string `json:"de"`
		PtBr []string `json:"pt-br"`
		Uk []string `json:"uk"`
		Cs []string `json:"cs"`
		Tr []string `json:"tr"`
		Ru []string `json:"ru"`
		It []string `json:"it"`
		Eo []string `json:"eo"`
		Bg []string `json:"bg"`
		Hi []string `json:"hi"`
		El []string `json:"el"`
		Ar []string `json:"ar"`
		Fr []string `json:"fr"`
		Hr []string `json:"hr"`
		Be []string `json:"be"`
		Hu []string `json:"hu"`
		En []string `json:"en"`
		Vi []string `json:"vi"`
		Nl []string `json:"nl"`
		Th []string `json:"th"`
		Ja []string `json:"ja"`
		Mr []string `json:"mr"`
	} `json:"transcript"`
}

func main() {
	


	
}
