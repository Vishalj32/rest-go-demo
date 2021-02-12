package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-go-demo/entity"
	"rest-go-demo/repository"

	"github.com/gorilla/mux"
)

//CreatePerson createPerson Handler
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(reqBody, &person)

	_, err := repository.CreatePerson(person)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(reqBody)
}

//GetPerson get Single Person handler
func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	person, err := repository.GetPerson(key)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No such data"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(person)
	}

}

//GetAllPersons GetAllData
func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := repository.GetAllPersons()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No such data"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

//DeletePerson removes person data
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	err := repository.DeletePerson(key)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No such data"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}

//UpdatePerson updates person data
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(reqBody, &person)

	res, err := repository.UpdatePerson(person)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

}
