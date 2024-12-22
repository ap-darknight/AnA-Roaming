package token_handling_strategies

import (
	context_objects "AnA-Roaming/ana-authenticator/dto-layer/context-objects"
	user_details_repo "AnA-Roaming/ana-authenticator/repository-layer/user-details-repo"
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtTokenHandler struct {
	Config *config_dto.Config
}

// Custom claims struct embedding jwt.RegisteredClaims and adding UserContext
type CustomClaims struct {
	context_objects.UserContext
	jwt.RegisteredClaims
}

func NewJwtTokenHandler(config *config_dto.Config) *JwtTokenHandler {
	return &JwtTokenHandler{
		Config: config,
	}
}

// GenerateToken method that takes a UserContext and creates a JWT
func (j JwtTokenHandler) GenerateToken(ctx context_objects.RepoContext, user user_details_repo.User) (string, error) {
	// Define token expiration time (e.g., 24 hours)
	expirationTime := time.Now().Add(time.Duration(j.Config.AuthKeys.JwtSecret.TokenExpirationTime * 24))

	// Create custom claims with UserContext and standard registered claims (expiration)
	claims := &CustomClaims{
		UserContext: context_objects.UserContext{
			Username: user.UserName,
			UserId:   user.UniqueID.String(),
			Role:     user.Role,
			Email:    user.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.Config.AuthKeys.JwtSecret.EncryptionKey))
	if err != nil {
		ctx.Logger.Errorw("Error signing token", "error", err)
		return "", err
	}

	return tokenString, nil
}

// ValidateToken method that parses and validates the JWT, returning the UserContext
func (j JwtTokenHandler) ValidateToken(ctx context_objects.RepoContext, tokenString string) (bool, error) {
	// Parse the token with the secret key
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.Config.AuthKeys.JwtSecret.EncryptionKey), nil
	})

	// Check if there was an error or the token is invalid
	if err != nil {
		ctx.Logger.Errorw("Error parsing token", "error", err)
		return false, err
	}

	// Validate the token claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		ctx.UserContext = claims.UserContext
		return true, nil
	}

	return true, jwt.ErrSignatureInvalid
}
