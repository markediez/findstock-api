package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

type Supply struct {
	Id uuid.UUID `json:"id"`
	Item string `json:"item"`
	Store string `json:"store"`
	Location string `json:"location"`
	CreatedAt int64 `json:"createdAt"`
}

var Supplies []Supply

func returnAllStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStocks")
	json.NewEncoder(w).Encode(Supplies)
}

func postNewStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: postNewStock")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var supply Supply
	json.Unmarshal(reqBody, &supply)
	supply.Id = uuid.Must(uuid.NewRandom())
	supply.CreatedAt = time.Now().UTC().Unix() * 1000
	Supplies = append(Supplies, supply)

	json.NewEncoder(w).Encode(supply)
}

func handleRequests() {
	router := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/supplies", returnAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/supplies", postNewStock).Methods("POST")

	log.Println("Listening to port 7777")
	err := http.ListenAndServe(":7777", handlers.CORS(header, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Supplies = []Supply {
		Supply { Id: uuid.Must(uuid.NewRandom()), Item: "Toilet Paper", Store: "Sacramento Co-op", Location: "2820 R St, Sacramento, CA 95816", CreatedAt: time.Now().AddDate(0, 0, -1).UTC().Unix() * 1000 },
		Supply { Id: uuid.Must(uuid.NewRandom()), Item: "Toilet Paper", Store: "Whole Foods Market", Location: "4315 Arden Way, Sacramento, CA 95864", CreatedAt: time.Now().AddDate(0, 0, -3).UTC().Unix() * 1000 },
		Supply { Id: uuid.Must(uuid.NewRandom()), Item: "Thermometer", Store: "CVS", Location: "3338 Arden Way, Sacramento, CA 95825", CreatedAt: time.Now().UTC().Unix() * 1000 },
	}

	handleRequests()
}