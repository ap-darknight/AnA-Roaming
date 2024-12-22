package authenticator_service

import (
	auth_strategies "AnA-Roaming/ana-authenticator/services/authenticator_service/auth-strategies"
	"go.uber.org/fx"
)

func NewAuthService() fx.Option {
	return fx.Provide(
		auth_strategies.NewAuthSimpleStrategy,
		auth_strategies.NewAuthEmailOtpStrategy,
		auth_strategies.NewAuthMultiFactorStrategy,
		auth_strategies.NewAuthJWTStrategy,
		auth_strategies.NewAuthMobileOTPStrategy,

		// Third Party Authenticators
	)
}
