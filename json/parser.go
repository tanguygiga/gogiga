package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"gogiga/log"
)

func Parser(path string, v interface{}) {

	log.Init(os.Stdout, os.Stdout)
	logError := log.Error.Println


	file, err := os.Open(path)
	if err != nil {
		logError("Error opening JSON file:", err)
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logError("Error reading JSON data:", err)
	}

	json.Unmarshal(data, &v)
}
