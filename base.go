package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(people)
}

func GetPersonEndpoint(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Person{})
}

func CreatePersonEndpoint(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(res).Encode(people)
}

func DeletePersonEndpoint(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dubling", State: "California"}})
	people = append(people, Person{ID: "2", FirstName: "Juan Carlos", LastName: "Ulloa", Address: &Address{City: "Santa Cruz de la Sierra", State: "Santa Cruz"}})
	people = append(people, Person{ID: "3", FirstName: "Don", LastName: "Cipiyos" /*, Address: &Address{City: "Cobija", State: "Pando"}*/})

	//Endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":3000", router))

}
