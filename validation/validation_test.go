package validation

import (
	"testing"
)

func TestValidateStateValidCases(t *testing.T) {
	validStates := []string{"open", "closed", "accepted", "investigating"}

	for _, state := range validStates {
		err := ValidateState(state)
		if err != nil {
			t.Errorf("expected no error for valid state: %s, got: %v", state, err)
		}
	}
}

func TestValidateStateInvalidCases(t *testing.T) {
	invalidStates := []string{"invalid", "pending", "on hold", "resolved", ""}

	for _, state := range invalidStates {
		err := ValidateState(state)
		if err == nil {
			t.Errorf("expected error for invalid state: %s, but got none", state)
		}
	}
}

func TestValidateStateEmptyState(t *testing.T) {
	err := ValidateState("")
	if err == nil || err.Error() != "state is required" {
		t.Errorf("expected error: 'state is required', got: %v", err)
	}
}
