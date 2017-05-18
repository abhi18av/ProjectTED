package main

import (
	"encoding/json"
	"fmt"
)

type UCSUserAccess struct {
	Groups                 map[string]string  `json:"groups"`
	Access                 map[string]string  `json:"access"`
	IsExpert               string             `json:"isExpert"`
	SubscriptionGroupDates map[string]GroupID `json:"subscription_group_dates"`
}

type GroupID struct {
	StartDate int `json:"START_DATE"`
	EndDate   int `json:"END_DATE"`
}

func main() {
	var test = UCSUserAccess{
		Access:                 map[string]string{"STOCK": "1", "FOREX": "1", "WEBFOREX": "1", "WEBSTOCK": "1"},
		SubscriptionGroupDates: map[string]GroupID{"32": GroupID{1464753600, 1472616000}, "42": GroupID{1470024000, 1472616000}},
	}
	body, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
