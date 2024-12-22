package auth_strategies

import (
	authenticator_objects2 "AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	mongo_infra "AnA-Roaming/repo-infra/mongo-infra"
	"log"
)

type AuthJWTStrategy struct {
	MongoDB mongo_infra.MongoInfra
}

func NewAuthJWTStrategy(mongoDB mongo_infra.MongoInfra) *AuthJWTStrategy {
	return &AuthJWTStrategy{
		MongoDB: mongoDB,
	}
}

func (a *AuthJWTStrategy) AuthenticateUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	log.Println("Authenticating user with JWT", a.MongoDB.GetSlaveDB())
	return authenticator_objects2.AuthResponse{}, nil
}

func (a *AuthJWTStrategy) AuthorizeUser(request authenticator_objects2.AuthRequest) (authenticator_objects2.AuthResponse, error) {
	return authenticator_objects2.AuthResponse{}, nil
}
