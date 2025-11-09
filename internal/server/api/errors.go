package api

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

var InvalidInputResponse = ErrorResponse{ErrorMessage: "Invalid input data"}
