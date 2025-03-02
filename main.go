package main

import (
	"bechallenge/handlers"
	"bechallenge/services/dataservice"
	"bechallenge/services/processingservice"
	"bechallenge/services/referralservice"
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
)

func main() {
	// Resolve dependencies and inject them where needed
	dataService, err := dataservice.NewDataService("data")
	if err != nil {
		panic(err.Error())
	}
	processingService := processingservice.NewProcessingService(&dataService)
	referralService := referralservice.NewReferralService(&dataService)
	userHandler := handlers.NewUserHandler(&processingService)
	userActionCountHandler := handlers.NewUserActionCountHandler(&processingService)
	nextActionsHandler := handlers.NewNextActionsHandler(&processingService)
	referralIndexesHandler := handlers.NewReferralIndexesHandler(&referralService)

	// Let program args define port
	flag.IntVar(&port, "port", 3000, "port to run the service on")
	flag.Parse()

	// Set up handlers and attach to routes
	mux := http.NewServeMux()
	mux.HandleFunc("/users/{id}", userHandler.Handle)
	mux.HandleFunc("/users/{id}/actions/count", userActionCountHandler.Handle)
	mux.HandleFunc("/actions/{type}/nextactions", nextActionsHandler.Handle)
	mux.HandleFunc("/referralindexes", referralIndexesHandler.Handle)

	// Start listening on that address and serving responses based on handlers above
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, mux)
}
