package postman

import "fmt"

type errorWrapper struct {
	Error Error `json:"error"`
}

// Error represents an error from the postman api.
type Error struct {
	Name    string            `json:"name"`
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

// Error satisfies the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}
