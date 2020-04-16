package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Post struct {
	Stock string `json:"Stock"`
	Store string `json:"Store"`
	Location string `json:"Location"`
}

var Posts []Post

func returnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPosts")
	json.NewEncoder(w).Encode(Posts)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/posts", returnAllPosts)

	log.Fatal(http.ListenAndServe(":7777", router))
}

func main() {
	Posts = []Post {
		Post { Stock: "Toilet Paper", Store: "Sacramento Co-op", Location: "2820 R St, Sacramento, CA 95816" },
		Post { Stock: "Toilet Paper", Store: "Whole Foods Market", Location: "4315 Arden Way, Sacramento, CA 95864" },
		Post { Stock: "Thermometer", Store: "CVS", Location: "3338 Arden Way, Sacramento, CA 95825" },
	}

	handleRequests()
}