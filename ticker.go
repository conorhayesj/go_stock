package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

func getTickers() (tickers[]string) {

	f, err := os.Open("tickers")
	if err != nil {
		log.Fatal(err)
	}

	//var tickers []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tickers = append(tickers, scanner.Text())
	}
	return tickers
}

func addTicker(ticker string) {

	tickers := getTickers()
	for _, b := range tickers {
		if b == ticker {
			fmt.Println(ticker, " is already entered")
			break
		}
	}

	f, err := os.OpenFile("tickers", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(ticker)
	if err != nil {
		log.Fatal(err)
	}
	f.Sync()

}
