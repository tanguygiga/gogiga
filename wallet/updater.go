package main

import (
	"gogiga/json"
	"gogiga/log"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/csv"
	"strings"
	"io"
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
			pricePerShare := 15.37
			price := float64(a.Line[i].Quantity) * pricePerShare
			logInfo("Ligne :", l.Symbol+"."+l.Market, "avec", l.Quantity, "titres :", price, "€")
			total += price
		}

	}
	logInfo("Votre portefeuille vaut aujourd'hui :", total, "€")

	base := "http://finance.yahoo.com/d/quotes.csv?"
	s := "s="
	f := "f="
	param := "sd1l1"
	url := base + s + "ALP.PA" + "&" + f + param

	resp, err := http.Get(url)
	if err != nil {
		logError("HTTP get not working")
	}

	logInfo(url)
	logInfo(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logInfo(body)
	in := `first_name,last_name,username
	Rob,Pike,rob
	Ken,Thompson,ken
	Robert,Griesemer,gri`

	r := csv.NewReader(strings.NewReader(in))

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
