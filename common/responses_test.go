package common

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	rr := httptest.NewRecorder()

	payload := map[string]string{"message": "test successful"}
	RespondWithJSON(rr, http.StatusOK, payload)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if response["message"] != "test successful" {
		t.Errorf("expected response message 'test successful', got %v", response["message"])
	}
}

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()

	errorMessage := "something went wrong"
	RespondWithError(rr, http.StatusBadRequest, errorMessage)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}

	var response JSONResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if response.Success != false {
		t.Errorf("expected Success false, got %v", response.Success)
	}

	if response.Error != errorMessage {
		t.Errorf("expected error message '%s', got '%s'", errorMessage, response.Error)
	}
}
