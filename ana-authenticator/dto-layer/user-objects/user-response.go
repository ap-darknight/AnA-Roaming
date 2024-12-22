package user_objects

import (
	"AnA-Roaming/ana-authenticator/dto-layer/error-objects"
)

// UserResponse is a struct that contains the response parameters for the User APIs.
type UserResponse struct {
	error_objects.BaseError
	Username string `json:"username,omitempty"`
	Role     string `json:"role,omitempty"`
}
