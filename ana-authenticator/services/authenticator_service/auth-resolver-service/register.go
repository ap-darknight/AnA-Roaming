package auth_resolver_service

import "go.uber.org/fx"

func NewAuthResolverService() fx.Option {
	return fx.Provide(
		NewAuthResolver,
	)
}
