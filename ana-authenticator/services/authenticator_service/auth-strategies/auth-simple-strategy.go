package auth_strategies

import (
	authenticator_objects2 "AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	mongo_infra "AnA-Roaming/repo-infra/mongo-infra"
)

type AuthSimpleStrategy struct {
	mongoDB mongo_infra.MongoInfra
}

func NewAuthSimpleStrategy(mongoDB mongo_infra.MongoInfra) *AuthSimpleStrategy {
	return &AuthSimpleStrategy{
		mongoDB: mongoDB,
	}
}

func (a *AuthSimpleStrategy) AuthenticateUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}

func (a *AuthSimpleStrategy) AuthorizeUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}
