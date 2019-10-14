package gorest

import (
	"encoding/json"
	"log"
	"net/http"

	"context"
)

// Error struct for API error
type Error struct {
	Message string `json:"message"`
}

// Response struct for API response
type Response struct {
	Status int         `json:"status"`
	Error  *Error      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// ReturnAPIError creates API error response
func ReturnAPIError(ctx context.Context, w http.ResponseWriter, err error) error {
	var response Response
	response.Status = 1
	response.Error = &Error{
		Message: err.Error(),
	}
	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		log.Printf("Error json err: %v, error: %v", jsonErr, err)
	}
	return err
}

// ReturnAPIResponse creates API response
func ReturnAPIResponse(ctx context.Context, w http.ResponseWriter, data interface{}) error {
	var response Response
	response.Status = 0
	response.Error = nil
	response.Data = data
	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		log.Printf("Response json err: %v, data: %#v", jsonErr, data)
	}
	return jsonErr
}
