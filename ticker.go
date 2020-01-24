package main

import (
	"os"
	"bufio"
	"log"
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

