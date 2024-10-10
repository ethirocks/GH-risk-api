package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
)

func TestCreateRiskSuccess(t *testing.T) {
	// Reset the global storage
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	payload := `{"title": "Test Risk", "description": "Risk description", "state": "open"}`
	req, err := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRisk)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, status)
	}

	var response common.Risk
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if response.Title != "Test Risk" {
		t.Errorf("expected Title 'Test Risk', got '%s'", response.Title)
	}

	if response.Description != "Risk description" {
		t.Errorf("expected Description 'Risk description', got '%s'", response.Description)
	}

	if response.State != "open" {
		t.Errorf("expected State 'open', got '%s'", response.State)
	}
}

func TestCreateRiskMissingFields(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	payload := `{"title": "Test Risk"}`
	req, err := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRisk)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}
}

func TestCreateRiskInvalidState(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	payload := `{"title": "Test Risk", "description": "Risk description", "state": "invalid-state"}`
	req, err := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRisk)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}
}
