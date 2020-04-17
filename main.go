package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Stock struct {
	Id int `json:"id"`
	Item string `json:"item"`
	Store string `json:"store"`
	Location string `json:"location"`
	CreatedAt int64 `json:"createdAt"`
}

var Stocks []Stock

func returnAllStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStocks")
	json.NewEncoder(w).Encode(Stocks)
}

func handleRequests() {
	router := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/stocks", returnAllStocks).Methods("GET", "OPTIONS")

	log.Println("Listening to port 7777")
	err := http.ListenAndServe(":7777", handlers.CORS(header, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Stocks = []Stock {
		Stock { Id: 1, Item: "Toilet Paper", Store: "Sacramento Co-op", Location: "2820 R St, Sacramento, CA 95816", CreatedAt: time.Now().AddDate(0, 0, -1).UTC().Unix() * 1000 },
		Stock { Id: 2, Item: "Toilet Paper", Store: "Whole Foods Market", Location: "4315 Arden Way, Sacramento, CA 95864", CreatedAt: time.Now().AddDate(0, 0, -3).UTC().Unix() * 1000 },
		Stock { Id: 3, Item: "Thermometer", Store: "CVS", Location: "3338 Arden Way, Sacramento, CA 95825", CreatedAt: time.Now().UTC().Unix() * 1000 },
	}

	handleRequests()
}