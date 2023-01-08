package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// elastic search is a backup solutioin
	// "github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

// func getElasticSearchClient() {
// 	client, err := elasticsearch.NewDefaultClient()
// 	log.Println(elasticsearch.Version)
// 	return client, err
// }

type payload struct {
	id          string
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

type payloadSearchParam struct {
	Title             string
	Version           string
	MaintainersEmails []string
	MaintainersNames  []string
	Company           string
	Website           string
	Source            string
	License           string
	Description       string
}

type allpayload []payload

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/payload", createPayload).Methods("PUT")
	router.HandleFunc("/payload/{id}", changePayload).Methods("POST")
	router.HandleFunc("/payload/{id}", deletePayload).Methods("DELETE")
	router.HandleFunc("/payload/{id}", getPayloadById).Methods("GET")
	router.HandleFunc("/payloads", getPayloadsByParams).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPayloadById(w http.ResponseWriter, r *http.Request) {
	payloadID := mux.Vars(r)["id"]
	returnPayload, err := fetchPayload(payloadID)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(returnPayload)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

}

func changePayload(w http.ResponseWriter, r *http.Request) {
	var newpayload payload
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter request data")
	}
	payloadID := mux.Vars(r)["id"]
	if payloadID != "" {
		fmt.Fprintf(w, "please input id")
	}

	yaml.Unmarshal(reqBody, &newpayload)

	validateResult, err := validateReq(newpayload)

	if validateResult { // valid input
		log.Println("input yaml is valid!")
		newpayload.id = payloadID
		SavePayload(newpayload)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newpayload)

	} else { //invalid input
		log.Println("input yaml is invalid!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
}

func deletePayload(w http.ResponseWriter, r *http.Request) {

	Id := mux.Vars(r)["id"]
	if Id == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("eventId is empty")
	}
	var isdeleted = deleteRecord(Id)
	if isdeleted {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("delete succeeded")
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("eventId doesn't exsit")
	}

}

func createPayload(w http.ResponseWriter, r *http.Request) {

	var newpayload payload
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter request data")
	}
	yaml.Unmarshal(reqBody, &newpayload)

	validateResult, err := validateReq(newpayload)

	if validateResult { // valid input
		log.Println("input yaml is valid!")
		var id string
		if !isPayloadExist(newpayload) {

			id = getUUID()
			newpayload.id = id
			log.Println("event does not exsit, saving...")
			SavePayload(newpayload)
			log.Println("event saved")

			w.WriteHeader(http.StatusCreated)
			fmt.Println("payload saved, payloadId : %s", id)
			json.NewEncoder(w).Encode(fmt.Sprintf("payload saved, payloadId : %s", id))
		} else {
			log.Println("input yaml payload exsiting already!")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(fmt.Sprintf("payload exsits already"))
		}
	} else { //invalid input
		log.Println("input yaml is invalid!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
}

// Data dedup, check if event exists in DB.
func isPayloadExist(newPayload payload) bool {
	var newPayloadSearchParam = FlatenPayloadSearchParam(newPayload)
	var resultList = searchPayloadByField(newPayloadSearchParam)
	return len(resultList) != 0
}

func getPayloadsByParams(w http.ResponseWriter, r *http.Request) {

	var searchParams payloadSearchParam

	if r.URL.Query().Get("Title") != "" {
		searchParams.Title = r.URL.Query().Get("Title")
	}

	if r.URL.Query().Get("Version") != "" {
		searchParams.Version = r.URL.Query().Get("Version")
	}

	if r.URL.Query().Get("MaintainersEmail") != "" {
		r.ParseForm()
		searchParams.MaintainersEmails = r.Form["MaintainersEmail"]
	}

	if r.URL.Query().Get("MaintainersName") != "" {
		r.ParseForm()
		searchParams.MaintainersNames = r.Form["MaintainersName"]
	}

	if r.URL.Query().Get("Company") != "" {
		searchParams.Company = r.URL.Query().Get("Company")
	}

	if r.URL.Query().Get("Website") != "" {
		searchParams.Website = r.URL.Query().Get("Website")
	}

	if r.URL.Query().Get("Source") != "" {
		searchParams.Source = r.URL.Query().Get("Source")
	}

	if r.URL.Query().Get("License") != "" {
		searchParams.License = r.URL.Query().Get("License")
	}

	if r.URL.Query().Get("Description") != "" {
		searchParams.Description = r.URL.Query().Get("Description")
	}

	var payloadIds = searchPayloadByField(searchParams)
	var eventsList = allpayload{}
	for _, id := range payloadIds {
		singlePayload, err := fetchPayload(id)
		if err != nil {
			//do nothing
		}
		eventsList = append(eventsList, singlePayload)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventsList)
}
