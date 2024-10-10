// store common structs
package common

// Risk represents the structure of a risk.
type Risk struct {
	ID          string `json:"id"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
