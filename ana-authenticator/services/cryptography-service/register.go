package cryptography_service

import (
	symmetric_strategies "AnA-Roaming/ana-authenticator/services/cryptography-service/symmetric-strategies"
	"go.uber.org/fx"
)

func NewCryptographyService() fx.Option {
	return fx.Provide(
		symmetric_strategies.NewChaCha20CryptographyService,
		symmetric_strategies.NewAESCryptographyService,
	)
}
