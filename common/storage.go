// store common structs and functions
package common

import (
	"fmt"
	"sync"
)

// RiskStorage holds risks in memory with fast lookup and insertion order tracking
type RiskStorage struct {
	Risks map[string]Risk // Map for fast lookups
	Order []string        // Slice to store risk IDs in order of insertion
	mu    sync.Mutex      // Mutex to handle concurrent access
}

// Global instance of the in-memory risk storage
var Storage = RiskStorage{
	Risks: make(map[string]Risk),
	Order: []string{},
}

// AddRisk adds a new risk to the storage
func (rs *RiskStorage) AddRisk(risk Risk) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if _, exists := rs.Risks[risk.ID]; exists {
		return fmt.Errorf("risk with ID %s already exists", risk.ID)
	}

	rs.Risks[risk.ID] = risk
	rs.Order = append(rs.Order, risk.ID)

	return nil
}

// GetAllRisks returns all risks in insertion order
func (rs *RiskStorage) GetAllRisks() ([]Risk, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if len(rs.Risks) == 0 {
		return nil, fmt.Errorf("no risks found")
	}

	riskList := make([]Risk, 0, len(rs.Risks))
	for _, riskID := range rs.Order {
		riskList = append(riskList, rs.Risks[riskID])
	}
	return riskList, nil
}

// GetRiskByID retrieves a specific risk by ID
func (rs *RiskStorage) GetRiskByID(id string) (Risk, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	risk, exists := rs.Risks[id]
	if !exists {
		return Risk{}, fmt.Errorf("risk with ID %s not found", id)
	}
	return risk, nil
}

// UpdateRisk updates an existing risk in the storage
func (rs *RiskStorage) UpdateRisk(id string, risk Risk) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	rs.Risks[id] = risk
}
