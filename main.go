package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Post struct {
	Id int `json:"id"`
	Stock string `json:"stock"`
	Store string `json:"store"`
	Location string `json:"location"`
}

var Posts []Post

func returnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPosts")
	json.NewEncoder(w).Encode(Posts)
}

func handleRequests() {
	router := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/posts", returnAllPosts).Methods("GET", "OPTIONS")

	log.Println("Listening to port 7777")
	err := http.ListenAndServe(":7777", handlers.CORS(header, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Posts = []Post {
		Post { Id: 1, Stock: "Toilet Paper", Store: "Sacramento Co-op", Location: "2820 R St, Sacramento, CA 95816" },
		Post { Id: 2, Stock: "Toilet Paper", Store: "Whole Foods Market", Location: "4315 Arden Way, Sacramento, CA 95864" },
		Post { Id: 3, Stock: "Thermometer", Store: "CVS", Location: "3338 Arden Way, Sacramento, CA 95825" },
	}

	handleRequests()
}