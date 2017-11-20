package yahoofinance

import (
	"encoding/csv"
	"fmt"
	"github.com/aktau/gofinance/fquery"
	"github.com/aktau/gofinance/util"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	QuotesUrl = "http://download.finance.yahoo.com/d/quotes.csv"
)

func csvQuotes(symbols []string) ([]fquery.Quote, error) {
	v := url.Values{}

	/* which symbols? */
	v.Set("s", strings.Join(symbols, ","))

	v.Set("f", "nsxabpoghkjl1m3m4ydr1")

	req, err := Csv(QuotesUrl, v)
	if err != nil {
		return nil, err
	}
	defer req.Close()
	r := csv.NewReader(req)

	results := make([]fquery.Quote, 0, len(symbols))
	for {
		fields, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		res := fquery.Quote{
			Name:             fields[0],
			Symbol:           fields[1],
			Ask:              floatOr(fields[3]),
			Bid:              floatOr(fields[4]),
			Open:             floatOr(fields[6]),
			PreviousClose:    floatOr(fields[5]),
			LastTradePrice:   floatOr(fields[11]),
			DayLow:           floatOr(fields[7]),
			DayHigh:          floatOr(fields[8]),
			YearLow:          floatOr(fields[9]),
			YearHigh:         floatOr(fields[10]),
			Ma50:             floatOr(fields[12]),
			Ma200:            floatOr(fields[13]),
			DividendYield:    floatOr(fields[14]),
			DividendPerShare: floatOr(fields[15]),
		}

		tm, err := time.Parse(fields[16], util.FmtMonthDay)
		if err == nil {
			res.DividendExDate = tm
		}

		results = append(results, res)
	}

	return results, nil
}

func floatOr(orig string) float64 {
	if f, err := strconv.ParseFloat(orig, 64); err == nil {
		return f
	} else {
		return 0
	}
}

func Csv(baseUrl string, params url.Values) (io.ReadCloser, error) {
	params.Set("e", ".csv")

	url := baseUrl + "?" + params.Encode()
	fmt.Println("csv: firing HTTP GET at ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

    © 2017 GitHub, Inc.
    Terms
    Privacy
    Security
    Status
    Help

    Contact GitHub
    API
    Training
    Shop
    Blog
    About


