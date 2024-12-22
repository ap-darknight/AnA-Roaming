package auth_strategies

import (
	authenticator_objects2 "AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	mongo_infra "AnA-Roaming/repo-infra/mongo-infra"
)

type AuthMultiFactorStrategy struct {
	mongoDB mongo_infra.MongoInfra
}

func NewAuthMultiFactorStrategy(mongoDB mongo_infra.MongoInfra) *AuthMultiFactorStrategy {
	return &AuthMultiFactorStrategy{
		mongoDB: mongoDB,
	}
}

func (a *AuthMultiFactorStrategy) AuthenticateUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}

func (a *AuthMultiFactorStrategy) AuthorizeUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}
