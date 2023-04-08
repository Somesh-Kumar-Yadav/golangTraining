package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`              // remove password
	Tags     []string `json:"tags,omitempty"` // omitempty if nil then so do not throw in json
}

func main() {
	fmt.Println("Json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React Bootcamp", 199, "Lco.in", "anc123", []string{"abc", "dd"}},
		{"Full Bootcamp", 399, "Lco.in", "ahk123", []string{"dkd", "iod"}},
		{"Backend Bootcamp", 499, "Lco.in", "huc123", nil},
	}

	// package this data as JSON data

	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
                "coursename": "Full Bootcamp",
                "price": 399,
                "platform": "Lco.in",
                "tags": [
                        "dkd",
                        "iod"
                ]
        } 
	`)
	var lcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v \n", lcoCourses)
	}

	// some cases where you just wnat to add data to key value

	var myOnlineMap map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineMap)

	fmt.Printf("%#v \n", myOnlineMap)

}
