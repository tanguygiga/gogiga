package json

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type Account struct {
	Cash float64 `json:"cash"`
	Line []Line  `json:"line"`
}

type Line struct {
	Symbol   string `json:"symbol"`
	Market   string `json:"market"`
	Quantity int    `json:"quantity"`
}

func accountGiver (jsonLocation String) Account {
	jsonFile, err := os.Open(jsonLocation)
	if err != nil {
		Error.Println("Error opening JSON file:", err)
	}

	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		Error.Println("Error reading JSON data:", err)
	}

	var a Account
	json.Unmarshal(data, &a)

	return a
}
