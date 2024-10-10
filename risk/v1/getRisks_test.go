package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
)

func TestGetRisksSuccess(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	common.Storage.AddRisk(common.Risk{ID: "1", Title: "Test Risk 1", Description: "Description 1", State: "open"})
	common.Storage.AddRisk(common.Risk{ID: "2", Title: "Test Risk 2", Description: "Description 2", State: "closed"})

	req, err := http.NewRequest("GET", "/v1/risks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRisks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	if rr.Body.Len() == 0 {
		t.Errorf("expected non-empty response body")
	}
}

func TestGetRisksNoContent(t *testing.T) {
	common.Storage = common.RiskStorage{
		Risks: make(map[string]common.Risk),
		Order: []string{},
	}

	req, err := http.NewRequest("GET", "/v1/risks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRisks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNoContent, status)
	}
}
