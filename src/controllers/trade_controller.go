package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"apex-trader-api/src/models"
)

func GetTrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Fetch trade from database
	trade := models.Trade{ID: id, Symbol: "AAPL", Price: 100.0, Quantity: 10, Timestamp: time.Now()}
	json.NewEncoder(w).Encode(trade)
}