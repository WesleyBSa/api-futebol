package main

import (
	"log"
	"net/http"

	"github.com/WesleyBSa/api-futebol/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tabela", controllers.GetTabela).Methods("GET")

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
