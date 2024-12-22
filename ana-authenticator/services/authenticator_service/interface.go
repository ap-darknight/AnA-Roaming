package authenticator_service

import (
	authenticator_objects2 "AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
)

type AuthService interface {
	// AuthenticateUser is a method that authenticates a user based on the request
	AuthenticateUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error)
	// AuthorizeUser is a method that authorizes a user based on the request based on system roles
	AuthorizeUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error)
}
