package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ContractHandler(httpWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	httpWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(httpWriter, "ContractAddress %v\n", vars["contractAddress"])
}
