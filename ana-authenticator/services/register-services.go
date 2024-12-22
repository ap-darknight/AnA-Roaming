package services

import (
	"AnA-Roaming/ana-authenticator/services/authenticator_service"
	"AnA-Roaming/ana-authenticator/services/authenticator_service/auth-resolver-service"
	cryptography_service "AnA-Roaming/ana-authenticator/services/cryptography-service"
	token_handling_service "AnA-Roaming/ana-authenticator/services/token-handling-service"
	"go.uber.org/fx"
)

func NewServices() fx.Option {
	return fx.Options(
		authenticator_service.NewAuthService(),
		auth_resolver_service.NewAuthResolverService(),
		token_handling_service.NewTokenHandlingService(),
		cryptography_service.NewCryptographyService(),
	)
}
