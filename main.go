package main

import (
	"bechallenge/handlers"
	"bechallenge/services"
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
)

func main() {
	dataService, err := services.NewDataService("data")
	if err != nil {
		panic(err.Error())
	}
	processingService := services.NewProcessingService(&dataService)
	userHandler := handlers.NewUserHandler(&processingService)
	userActionCountHandler := handlers.NewUserActionCountHandler(&processingService)
	nextActionsHandler := handlers.NewNextActionsHandler(&processingService)

	flag.IntVar(&port, "port", 3000, "port to run the service on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/users/{id}", userHandler.Handle)
	mux.HandleFunc("/users/{id}/actions/count", userActionCountHandler.Handle)
	mux.HandleFunc("/actions/{type}/nextactions", nextActionsHandler.Handle)
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, mux)
}
