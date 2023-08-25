package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags, omitempty"`
}

func main() {
	println("MAKing of JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"React Bootcamp", 299, "LCO.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootamp", 199, "LCO.in", "abc123", []string{"web-dev", "Full-stack", "js"}},
		{"MongoDB", 129, "LCO.in", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourse, " ", "\t")

	checkNilError(err)
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootamp",
		"Price": 199,
		"tags": ["web-dev","Full-stack","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key:%v,Value:%v & type:%T", k, v, v)
	}

}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
