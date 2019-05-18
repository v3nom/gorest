package gorest

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
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
		log.Errorf(ctx, "Error json err: %v, error: %v", jsonErr, err)
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
		log.Errorf(ctx, "Response json err: %v, data: %#v", jsonErr, data)
	}
	return jsonErr
}
