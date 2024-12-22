package auth_resolver_service

import (
	"AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	"AnA-Roaming/ana-authenticator/services/authenticator_service"
)

type AuthResolver interface {
	AuthStrategyResolver(request authenticator_objects.AuthRequest) authenticator_service.AuthService
}
