package error_objects

// BaseError is a struct that contains the error parameters for the APIs.
type BaseError struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	StatusCode   string `json:"statusCode,omitempty"`
	Success      bool   `json:"success,omitempty"`
}
