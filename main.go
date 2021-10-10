package main

import (
	"net/http"

	"github.com/ashishbabar/go-eth-api-router/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", "4397")
}
func main() {
	zapLogger := getLogger()
	router := mux.NewRouter()
	port := viper.GetString("port")
	router.HandleFunc("/contract/{contractAddress}", handlers.ContractHandler(zapLogger))
	http.Handle("/", router)
	zapLogger.Info("Starting routes API at " + port)
	http.ListenAndServe(":"+port, router)
}
