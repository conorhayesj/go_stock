package main

import (
	"fmt"
	"net/http"
	//"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

const url = "https://query1.finance.yahoo.com/v7/finance/quote?symbols="

type Stock struct {

	Price	string	`json:"regularMarketPrice"`
	High	string	`json:"regularMarketDayHigh"`
	Low	string	`json:"regularMarketDayLow"`

}

func main() {

	tickers := []string{"AMD","ADI","AAPL","BMRN"}

	for _, ticker := range tickers {

		stockUrl := url + ticker
		r, err := http.Get(stockUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		bjson,err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var dat map[string]interface{}
		json.Unmarshal(bjson,&dat)

		bjson2, _ := json.Marshal(dat["quoteResponse"])
		var dat2 map[string][]map[string]interface{}
		json.Unmarshal(bjson2, &dat2)

		results := dat2["result"]

		quote := new(Stock)

		for _, raw := range results {
			result := map[string]string{}
			for i, k := range raw {
				result[i] = fmt.Sprintf("%v", k)
			}
			quote.Price = result["regularMarketPrice"]
			quote.High = result["regularMarketDayHigh"]
			quote.Low = result["regularMarketDayLow"]
		}

		fmt.Println(ticker, "STOCK\n-----------------")
		fmt.Println("Current Price:\t ", quote.Price)
		fmt.Println("Today's High:\t ", quote.High)
		fmt.Println("Today's Low:\t ", quote.Low, "\n")
	}
}
