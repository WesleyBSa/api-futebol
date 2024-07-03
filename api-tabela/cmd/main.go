package main

import (
	"log"
	"net/http"

	"github.com/WesleyBSa/api-tabela/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tabela", controllers.GetTabela).Methods("GET")

	log.Println("Listening on port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}
