package main

import (
	"log"
	"time"
)

func main() {
	mainLoop()

}

func mainLoop() {
	for {
		rates, err := fetchLatestExchangeRate()
		if err != nil {
			log.Println("Skipping updating exchange rate due to error")
			continue
		}
		sekRate := rates["SEK"]
		log.Println("latest rate", sekRate)

		err = publishToExchangeRate(ExchangeRateMessage{
			Currency: sekRate.Symbol,
			Rate:     sekRate.Last,
		})
		if err != nil {
			log.Println("Failed to publish to exchange", err)
		}

		time.Sleep(time.Minute)
	}
}
