package main


import (
	_"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func GetPeople(w http.ResponseWriter,r *http.Request) {}


func main(){

	router := mux.NewRouter()
	router.HandleFunc("/people",GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000",router))

}
