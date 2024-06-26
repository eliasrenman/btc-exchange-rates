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

		publishToExchangeRate(ExchangeRateMessage{
			Currency: sekRate.Symbol,
			Rate:     sekRate.Last,
		})

		time.Sleep(time.Minute)
	}
}
