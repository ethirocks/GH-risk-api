package common

import (
	"testing"
)

func createSampleRisk(id, state, title, description string) Risk {
	return Risk{
		ID:          id,
		State:       state,
		Title:       title,
		Description: description,
	}
}

func TestAddRisk(t *testing.T) {
	rs := &RiskStorage{
		Risks: make(map[string]Risk),
		Order: []string{},
	}

	risk := createSampleRisk("1", "open", "Test Risk", "Description of test risk")

	err := rs.AddRisk(risk)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	err = rs.AddRisk(risk)
	if err == nil || err.Error() != "risk with ID 1 already exists" {
		t.Errorf("expected error 'risk with ID 1 already exists', but got: %v", err)
	}
}

func TestGetAllRisks(t *testing.T) {
	rs := &RiskStorage{
		Risks: make(map[string]Risk),
		Order: []string{},
	}

	_, err := rs.GetAllRisks()
	if err == nil || err.Error() != "no risks found" {
		t.Errorf("expected error 'no risks found', but got: %v", err)
	}

	risk := createSampleRisk("1", "open", "Test Risk", "Description of test risk")
	rs.AddRisk(risk)

	risks, err := rs.GetAllRisks()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
	if len(risks) != 1 {
		t.Errorf("expected 1 risk, but got: %d", len(risks))
	}
	if risks[0].ID != "1" {
		t.Errorf("expected risk ID '1', but got: %s", risks[0].ID)
	}
}

func TestGetRiskByID(t *testing.T) {
	rs := &RiskStorage{
		Risks: make(map[string]Risk),
		Order: []string{},
	}

	_, err := rs.GetRiskByID("1")
	if err == nil || err.Error() != "risk with ID 1 not found" {
		t.Errorf("expected error 'risk with ID 1 not found', but got: %v", err)
	}

	risk := createSampleRisk("1", "open", "Test Risk", "Description of test risk")
	rs.AddRisk(risk)

	retrievedRisk, err := rs.GetRiskByID("1")
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
	if retrievedRisk.ID != "1" {
		t.Errorf("expected risk ID '1', but got: %s", retrievedRisk.ID)
	}
}

func TestUpdateRisk(t *testing.T) {
	rs := &RiskStorage{
		Risks: make(map[string]Risk),
		Order: []string{},
	}

	risk := createSampleRisk("1", "open", "Test Risk", "Description of test risk")
	rs.AddRisk(risk)

	updatedRisk := createSampleRisk("1", "closed", "Updated Risk Title", "Updated description")
	rs.UpdateRisk("1", updatedRisk)

	retrievedRisk, err := rs.GetRiskByID("1")
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
	if retrievedRisk.State != "closed" {
		t.Errorf("expected state 'closed', but got: %s", retrievedRisk.State)
	}
	if retrievedRisk.Title != "Updated Risk Title" {
		t.Errorf("expected title 'Updated Risk Title', but got: %s", retrievedRisk.Title)
	}
	if retrievedRisk.Description != "Updated description" {
		t.Errorf("expected description 'Updated description', but got: %s", retrievedRisk.Description)
	}
}
