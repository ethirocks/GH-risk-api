package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	v1 "github.com/ethirajmudhaliar/GH-risk-api/risk/v1"
	"github.com/gorilla/mux"
)

func TestLoggingMiddleware(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/v1/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	router.Use(LoggingMiddleware)

	req, err := http.NewRequest("GET", "/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	start := time.Now()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	if elapsed := time.Since(start); elapsed <= 0 {
		t.Errorf("expected a non-zero request processing time, got %v", elapsed)
	}
}

func TestRoutes(t *testing.T) {
	router := mux.NewRouter()

	router.HandleFunc("/v1/risks", v1.GetRisks).Methods("GET")
	router.HandleFunc("/v1/risks", v1.CreateRisk).Methods("POST")
	router.HandleFunc("/v1/risks/{id}", v1.GetRiskByID).Methods("GET")
	router.HandleFunc("/v1/risks/{id}", v1.UpdateRisk).Methods("PUT")

	req, err := http.NewRequest("GET", "/v1/risks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusNoContent {
		t.Errorf("expected status code %d or %d, got %d", http.StatusOK, http.StatusNoContent, status)
	}
}

func TestCreateRiskRoute(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/v1/risks", v1.CreateRisk).Methods("POST")

	req, err := http.NewRequest("POST", "/v1/risks", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest && status != http.StatusCreated {
		t.Errorf("expected status code %d or %d, got %d", http.StatusBadRequest, http.StatusCreated, status)
	}
}

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()

	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}
	common.Storage.AddRisk(common.Risk{ID: "1", Title: "Test Risk", Description: "Test Description", State: "open"})

	tests := []struct {
		method      string
		url         string
		body        []byte
		statusCode  int
		description string
	}{
		{"GET", "/v1/risks", nil, http.StatusOK, "Get risks"}, // Risks exist
		{"POST", "/v1/risks", []byte(`{"title": "New Risk", "description": "A test risk", "state": "open"}`), http.StatusCreated, "Create risk"},
		{"GET", "/v1/risks/non-existent-id", nil, http.StatusNoContent, "Get risk by ID not found"},
		{"PUT", "/v1/risks/1", []byte(`{"title": "Updated Title", "description": "Updated Description", "state": "closed"}`), http.StatusOK, "Update risk"},
	}

	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(tt.body))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()

		start := time.Now()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("%s: expected status code %d, got %d", tt.description, tt.statusCode, status)
		}

		if elapsed := time.Since(start); elapsed <= 0 {
			t.Errorf("expected a non-zero request processing time, got %v", elapsed)
		}
	}
}
