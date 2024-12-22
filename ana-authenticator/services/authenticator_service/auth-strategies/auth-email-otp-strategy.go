package auth_strategies

import (
	authenticator_objects2 "AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	mongo_infra "AnA-Roaming/repo-infra/mongo-infra"
)

type AuthEmailOtpStrategy struct {
	mongoDB mongo_infra.MongoInfra
}

func NewAuthEmailOtpStrategy(mongoDB mongo_infra.MongoInfra) *AuthEmailOtpStrategy {
	return &AuthEmailOtpStrategy{
		mongoDB: mongoDB,
	}
}

func (a *AuthEmailOtpStrategy) AuthenticateUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}

func (a *AuthEmailOtpStrategy) AuthorizeUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}
