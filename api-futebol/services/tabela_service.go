package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/WesleyBSa/api-futebol/models"
)

func FetchTabela() ([]models.Tabela, error) {
	log.Println("Fetching tabela from internal API")

	resp, err := http.Get("http://localhost:8001/tabela")
	if err != nil {
		log.Println("Error making GET request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		return nil, errors.New("failed to fetch tabela")
	}

	var tabela []models.Tabela
	if err := json.NewDecoder(resp.Body).Decode(&tabela); err != nil {
		log.Println("Error decoding JSON response:", err)
		return nil, err
	}

	log.Println("Successfully fetched tabela")
	return tabela, nil
}
