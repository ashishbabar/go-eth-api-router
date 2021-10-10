package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func ContractHandler(logger *zap.Logger) http.HandlerFunc {
	return func(httpWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		httpWriter.WriteHeader(http.StatusOK)
		logger.Info("Calling contract service with paratmeter " + vars["contractAddress"])
		fmt.Fprintf(httpWriter, "ContractAddress :%v\n", vars["contractAddress"])
	}

}
