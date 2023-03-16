package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Body       *responseBody
	StatusCode int
}
type responseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ToJSON writes the response to the given http.ResponseWriter
// with an application/json Content-Type header.
func (r Response) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r.Body)
}

// OK returns a new successful response.
func OK(message string, data interface{}) *Response {
	return newResponse(true, message, data, http.StatusOK)
}

// Fail returns a new failed response.
func Fail(message string, statusCode int) *Response {
	return newResponse(false, message, nil, statusCode)
}

func newResponse(success bool, message string, data interface{}, statusCode int) *Response {
	return &Response{
		Body: &responseBody{
			Success: success,
			Message: message,
			Data:    data,
		},
		StatusCode: statusCode,
	}
}
