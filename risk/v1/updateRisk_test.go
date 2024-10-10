package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	"github.com/gorilla/mux"
)

func TestUpdateRiskSuccess(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	existingRisk := common.Risk{ID: "1", Title: "Old Title", Description: "Old Description", State: "open"}
	common.Storage.AddRisk(existingRisk)

	updatedRisk := common.Risk{Title: "New Title", Description: "New Description", State: "closed"}
	payload, _ := json.Marshal(updatedRisk)

	req, err := http.NewRequest("PUT", "/v1/risks/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", UpdateRisk)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	var response common.Risk
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if response.Title != "New Title" {
		t.Errorf("expected Title 'New Title', got '%s'", response.Title)
	}

	if response.Description != "New Description" {
		t.Errorf("expected Description 'New Description', got '%s'", response.Description)
	}

	if response.State != "closed" {
		t.Errorf("expected State 'closed', got '%s'", response.State)
	}
}

func TestUpdateRiskNotFound(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	payload := `{"title": "New Title", "description": "New Description", "state": "closed"}`
	req, err := http.NewRequest("PUT", "/v1/risks/non-existent-id", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", UpdateRisk)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, status)
	}
}

func TestUpdateRiskInvalidPayload(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	existingRisk := common.Risk{ID: "1", Title: "Old Title", Description: "Old Description", State: "open"}
	common.Storage.AddRisk(existingRisk)

	payload := `{"title":`
	req, err := http.NewRequest("PUT", "/v1/risks/1", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", UpdateRisk)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRiskInvalidState(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	existingRisk := common.Risk{ID: "1", Title: "Old Title", Description: "Old Description", State: "open"}
	common.Storage.AddRisk(existingRisk)

	payload := `{"state": "invalid-state"}`
	req, err := http.NewRequest("PUT", "/v1/risks/1", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", UpdateRisk)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}
}
