package main

import (
	"encoding/json"
	"fmt"
)

type UCSUserAccess struct {
	Groups   map[string]string `json:"groups"`
	Access   map[string]string `json:"access"`
	IsExpert string            `json:"isExpert"`
}

func main() {
	var test = UCSUserAccess{
		Access: map[string]string{"STOCK": "1", "FOREX": "1", "WEBFOREX": "1", "WEBSTOCK": "1"},
	}
	body, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
