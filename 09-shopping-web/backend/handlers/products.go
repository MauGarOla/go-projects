package handlers

import (
	"encoding/json"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	json.NewEncoder(w).Encode(Products)
}
