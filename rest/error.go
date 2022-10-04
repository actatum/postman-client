// Package rest provides types/client for making requests to the postman REST api.
package rest

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ErrorResponse represents an error response from the postman api.
type ErrorResponse struct {
	Error *Error `json:"error"`
}

// Error represents the error fields from the postman api error response.
type Error struct {
	Name    string            `json:"name"`
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

// Error satisfies the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

// UnmarshalJSON customizes the json unmarshalling of Error.
func (e *Error) UnmarshalJSON(data []byte) error {
	var errorResult map[string]interface{}
	if err := json.Unmarshal(data, &errorResult); err != nil {
		return err
	}

	var name, message interface{}
	var ok bool

	name, ok = errorResult["name"]
	if !ok {
		return errors.New("invalid postman error response")
	}

	switch v := name.(type) {
	case string:
		e.Name = v
		message, ok = errorResult["message"]
		if !ok {
			return errors.New("missing postman error message")
		}
		e.Message, ok = message.(string)
		if !ok {
			return errors.New("error.message is not of type string")
		}
	case map[string]interface{}:
		name, ok = errorResult["name"].(map[string]interface{})["name"]
		if !ok {
			return errors.New("missing api security error name")
		}
		e.Name, ok = name.(string)
		if !ok {
			return errors.New("error.name.name is not of type string")
		}
		message, ok = errorResult["name"].(map[string]interface{})["reason"]
		if !ok {
			return errors.New("missing api security error reason")
		}
		e.Message, ok = message.(string)
		if !ok {
			return errors.New("error.name.reason is not of type string")
		}
	}

	return nil
}
