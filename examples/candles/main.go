package main

import (
	"log"

	"github.com/larscom/go-bitvavo/v2"
)

func main() {
	ws, err := bitvavo.NewWebSocket()
	if err != nil {
		log.Fatal(err)
	}

	candlechn, err := ws.Candles().Subscribe("ETH-EUR", "5m", 0)
	if err != nil {
		log.Fatal(err)
	}

	for candleEvent := range candlechn {
		log.Println(candleEvent)
	}
}
