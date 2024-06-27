package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func registerServer(port uint16) {
	log.Println(fmt.Sprintf("Server started on port: %d", port))
	http.HandleFunc("/exchange-rate/latest", getLatestExchangeRate)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln("Failed starting server!")
	}
}
func getLatestExchangeRate(w http.ResponseWriter, r *http.Request) {
	exchange, err := queryLatestExchange()
	if err != nil {
		log.Println("ERROR: Failed reading latest exchange rate")
		io.WriteString(w, `{"status":500, "message": "Server error"}`)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(exchange)
	if err != nil {
		log.Println("ERROR: Failed serializing latest exchange rate")
		io.WriteString(w, `{"status":500, "message": "Server error"}`)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
