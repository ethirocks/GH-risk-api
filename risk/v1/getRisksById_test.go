package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	"github.com/gorilla/mux"
)

func TestGetRiskByIDSuccess(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	common.Storage.AddRisk(common.Risk{ID: "1", Title: "Test Risk", Description: "Description of test risk", State: "open"})

	req, err := http.NewRequest("GET", "/v1/risks/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", GetRiskByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	if rr.Body.Len() == 0 {
		t.Errorf("expected non-empty response body")
	}
}

func TestGetRiskByIDNotFound(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	req, err := http.NewRequest("GET", "/v1/risks/non-existent-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", GetRiskByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNoContent, status)
	}
}
