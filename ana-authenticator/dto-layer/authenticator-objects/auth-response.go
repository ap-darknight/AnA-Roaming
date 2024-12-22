package authenticator_objects

import (
	"AnA-Roaming/ana-authenticator/dto-layer/error-objects"
)

type AuthResponse struct {
	error_objects.BaseError
	IsAuthenticated bool   `json:"is_authenticated,omitempty"` // Used for Password Verification
	IsAuthorized    bool   `json:"is_authorized,omitempty"`    // Used for Role Verification
	Token           string `json:"token,omitempty"`
}
