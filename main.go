package main

import (
	"net/http"

	"github.com/ashishbabar/go-eth-router/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contract/{contractAddress}", handlers.ContractHandler)
	http.Handle("/", router)

	http.ListenAndServe(":3486", router)
}