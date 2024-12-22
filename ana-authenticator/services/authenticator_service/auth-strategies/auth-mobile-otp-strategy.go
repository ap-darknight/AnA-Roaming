package auth_strategies

import (
	authenticator_objects2 "AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	mongo_infra "AnA-Roaming/repo-infra/mongo-infra"
)

type AuthMobileOTPStrategy struct {
	mongoDB mongo_infra.MongoInfra
}

func NewAuthMobileOTPStrategy(mongoDB mongo_infra.MongoInfra) *AuthMobileOTPStrategy {
	return &AuthMobileOTPStrategy{
		mongoDB: mongoDB,
	}
}

func (a *AuthMobileOTPStrategy) AuthenticateUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}

func (a *AuthMobileOTPStrategy) AuthorizeUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}
