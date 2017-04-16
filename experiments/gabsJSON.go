/*

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs"
)

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


*/