package main

import (
	"encoding/csv"
	"gogiga/json"
	"gogiga/log"
	"io"
	"net/http"
	"os"
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

func main() {

	log.Init(os.Stdout, os.Stdout)
	logInfo := log.Info.Println
	logError := log.Error.Println

	cd, _ := os.Getwd()
	path := cd + "/wallet/account.json"

	var a Account
	json.Parser(path, &a)

	total := a.Cash
	for i := 0; i < len(a.Line); i++ {
		l := a.Line[i]
		if l.Quantity > 0 {
			pricePerShare := 15.76
			price := float64(a.Line[i].Quantity) * pricePerShare
			logInfo("Ligne :", l.Symbol+"."+l.Market, "avec", l.Quantity, "titres :", price, "€")
			total += price
		}

	}
	logInfo("Votre portefeuille vaut aujourd'hui :", total, "€")

	//partie sur la récupération des infos de yahoo finance en csv
	//construction de l'url, récupération du csv, puis lecture du csv
	url := "http://finance.yahoo.com/d/quotes.csv?"
	s := "s="
	stocks := "ALP.PA+SOP.PA"
	s += stocks
	f := "&f="
	param := "sd1l1"
	f += param
	url += s + f

	resp, err := http.Get(url)
	if err != nil {
		logError("HTTP get not working")
	}

	logInfo(url)

	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logError(err)
		}

		logInfo(record)
	}
}
