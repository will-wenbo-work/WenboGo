package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// elastic search is a backup solutioin
	// "github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// func getESClient() {
// 	client, err := elasticsearch.NewDefaultClient()
// 	log.Println(elasticsearch.Version)
// 	return client, err

// }

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getEventsByParams).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type event struct {
	Title       string       `json:"Title"`
	Version     string       `json:"Version"`
	Maintainers []maintainer `json:"Maintainers"`
	Company     string       `json:"Company"`
	Website     string       `json:"Website"`
	Source      string       `json:"Source"`
	License     string       `json:"License"`
	Description string       `json:"Description"`
}

type maintainer struct {
	Name  string
	Email string
}

type allEvents []event

func createEvent(w http.ResponseWriter, r *http.Request) {

	var newEvent event
	// validator
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)

	if !isEventExist(newEvent) {
		SaveEvent(newEvent)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

// Data dedup, if event exist in DB, stop from save it again
func isEventExist(newEvent event) bool {
	var resultList = searchEventByField(newEvent)
	return len(resultList) != 0
}

func getEventsByParams(w http.ResponseWriter, r *http.Request) {

	var eventParams event
	if mux.Vars(r)["Title"] != "" {
		eventParams.Title = mux.Vars(r)["Title"]
	}

	var eventIds = searchEventByField(eventParams)
	eventIds = append(eventIds, 0)
}
