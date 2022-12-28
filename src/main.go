package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "./mockDB"
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
	router.HandleFunc("/events", getEvents).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type event struct {
	Title       string
	version     string
	maintainers []maintainer
	company     string
	website     string
	source      string
	license     string
	Description string `json:"Description"`
}

type maintainer struct {
	name  string
	email string
}

type allEvents []event

func createEvent(w http.ResponseWriter, r *http.Request) {
	// newEvent := event{
	// 	Title:       "title",
	// 	version:     "name",
	// 	maintainers: nil,
	// 	company:     "company",
	// 	website:     "web",
	// 	source:      "source",
	// 	license:     "license",
	// 	Description: "Description",
	// }

	var newEvent event
	// validator
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)

	SaveEvent(newEvent)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getEvents(w http.ResponseWriter, r *http.Request) {

	// eventID := mux.Vars(r)["id"]

	// for _, singleEvent := range events {
	// 	if singleEvent.ID == eventID {
	// 		json.NewEncoder(w).Encode(singleEvent)
	// 	}
	// }
}
