package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/WesleyBSa/api-futebol/services"
)

func GetTabela(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /tabela")
	tabela, err := services.FetchTabela()
	if err != nil {
		log.Println("Error fetching tabela:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tabela)
	log.Println("Successfully returned tabela")
}
