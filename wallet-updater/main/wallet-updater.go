package main

import (
	"os"
	"gogiga/wallet-updater/pkg/json"
	"gogiga/wallet-updater/pkg/logger"
)



func main() {
	logger.InitLog(os.Stdout, os.Stdout)

	cd, _ := os.Getwd()
	jsonLocation := cd + "/wallet-updater/account.json"

	a := json.accountGiver(jsonLocation)

	total := a.Cash
	for i := 0; i < len(a.Line); i++ {
		l := a.Line[i]
		if l.Quantity > 0 {
			pricePerShare := 15.37
			price := float64(a.Line[i].Quantity) * pricePerShare
			Info.Println("Ligne :", l.Symbol+"."+l.Market, "avec", l.Quantity, "titres :", price, "€")
			total += price
		}

	}
	Info.Println("Votre portefeuille vaut aujourd'hui :", total, "€")
}

type Account struct {
	Cash float64 `json:"cash"`
	Line []Line  `json:"line"`
}

type Line struct {
	Symbol   string `json:"symbol"`
	Market   string `json:"market"`
	Quantity int    `json:"quantity"`
}
