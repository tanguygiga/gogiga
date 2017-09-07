package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Je m'appelle Tanguy Gigarel et voici\nMon premier programme en go !\n")
	fmt.Printf("Ceci est mon adresse ip publique :\n")
	os.Stdout.Write(ip)
	fmt.Printf("\n")
}
