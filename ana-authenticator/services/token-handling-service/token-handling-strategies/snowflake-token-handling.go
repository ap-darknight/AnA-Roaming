package token_handling_strategies

import (
	context_objects "AnA-Roaming/ana-authenticator/dto-layer/context-objects"
	crypto_objects "AnA-Roaming/ana-authenticator/dto-layer/crypto-objects"
	user_details_repo "AnA-Roaming/ana-authenticator/repository-layer/user-details-repo"
	cryptography_service "AnA-Roaming/ana-authenticator/services/cryptography-service"
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"time"
)

type SnowflakeTokenHandler struct {
	CryptoService cryptography_service.CryptographyService
	Config        *config_dto.Config
}

func NewSnowflakeTokenHandler(cryptoService cryptography_service.CryptographyService, config *config_dto.Config) *SnowflakeTokenHandler {
	return &SnowflakeTokenHandler{
		CryptoService: cryptoService,
		Config:        config,
	}
}

func (s SnowflakeTokenHandler) GenerateToken(ctx context_objects.RepoContext, user user_details_repo.User) (string, error) {
	cryptoData := crypto_objects.CryptoObject{
		Key1: user.Email,
		Key2: user.PasswordHash,
		Key3: time.Now().Add(time.Duration(s.Config.AuthKeys.Snowflake.TokenExpirationTime * 24)).Unix(),
	}
	token, err := s.CryptoService.EncryptDataWithKey(cryptoData)
	if err != nil {
		ctx.Logger.Errorw("Error generating token", "error", err)
		return "", err
	}

	return token, nil
}

func (s SnowflakeTokenHandler) ValidateToken(ctx context_objects.RepoContext, token string) (bool, error) {
	decryptedData, err := s.CryptoService.DecryptDataWithKey(token)
	if err != nil {
		ctx.Logger.Errorw("Error validating token", "error", err)
		return false, err
	}

	cryptoData, ok := decryptedData.(crypto_objects.CryptoObject)
	if !ok {
		ctx.Logger.Errorw("Error validating token", "error", err)
		return false, err
	}

	if cryptoData.Key3 < time.Now().Unix() {
		// Token expired
		ctx.Logger.Errorw("Token Expired")
		return false, nil
	}

	return true, nil
}
