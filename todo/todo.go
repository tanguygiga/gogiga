package main

import (
	"fmt"
	"gogiga/json"
	"io/ioutil"
	"os"
)

type config struct {
	path string `json:"path"`
	Port int    `json:"port"`
}

func main() {
	wd, _ := os.Getwd()
	path := wd + "/config.json"

	var c config
	json.Parser(path, &c)

	b, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(b))
}
