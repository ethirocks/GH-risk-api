package main

import (
	"net/http"
	"time"

	"github.com/ethirajmudhaliar/GH-risk-api/logger"
	v1 "github.com/ethirajmudhaliar/GH-risk-api/risk/v1"
	"github.com/gorilla/mux"
)

// LoggingMiddleware logs details about incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Info("Started " + r.Method + " " + r.RequestURI)

		next.ServeHTTP(w, r)

		logger.LogRequest(r.Method, r.RequestURI, start)
	})
}

// SetupRouter sets up the router and middleware
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Define the routes for the API
	router.HandleFunc("/v1/risks", v1.GetRisks).Methods("GET")
	router.HandleFunc("/v1/risks", v1.CreateRisk).Methods("POST")
	router.HandleFunc("/v1/risks/{id}", v1.GetRiskByID).Methods("GET")
	router.HandleFunc("/v1/risks/{id}", v1.UpdateRisk).Methods("PUT")

	// Add the logging middleware
	router.Use(LoggingMiddleware)

	return router
}

func main() {
	router := SetupRouter()

	logger.Info("Starting server on port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Error("Error starting server: " + err.Error())
	}
}
