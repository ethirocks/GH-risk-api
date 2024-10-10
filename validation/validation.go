// validation
package validation

import "fmt"

// Define valid states
var validStates = map[string]bool{
	"open":          true,
	"closed":        true,
	"accepted":      true,
	"investigating": true,
}

// ValidateState checks if the provided state is valid
func ValidateState(state string) error {
	if state == "" {
		return fmt.Errorf("state is required")
	}
	if !validStates[state] {
		return fmt.Errorf("invalid state value: %s", state)
	}
	return nil
}
