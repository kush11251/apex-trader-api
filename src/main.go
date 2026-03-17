package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Apex Trader API started")
	r := mux.NewRouter()
	r.HandleFunc("/metrics", getMetrics).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Metrics endpoint"})
}