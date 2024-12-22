package token_handling_service

import (
	token_handling_strategies "AnA-Roaming/ana-authenticator/services/token-handling-service/token-handling-strategies"
	"go.uber.org/fx"
)

func NewTokenHandlingService() fx.Option {
	return fx.Provide(
		token_handling_strategies.NewJwtTokenHandler,
		token_handling_strategies.NewSnowflakeTokenHandler,
	)
}
