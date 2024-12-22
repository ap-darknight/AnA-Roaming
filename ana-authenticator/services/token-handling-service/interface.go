package token_handling_service

import (
	context_objects "AnA-Roaming/ana-authenticator/dto-layer/context-objects"
	user_details_repo "AnA-Roaming/ana-authenticator/repository-layer/user-details-repo"
)

type TokenHandlingService interface {
	GenerateToken(ctx context_objects.RepoContext, user user_details_repo.User) (string, error)
	ValidateToken(ctx context_objects.RepoContext, token string) (bool, error)
}
