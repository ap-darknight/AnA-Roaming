package auth_resolver_service

import (
	"AnA-Roaming/ana-authenticator/dto-layer/authenticator-objects"
	"AnA-Roaming/ana-authenticator/services/authenticator_service"
)

type AuthResolverImpl struct {
	JWTAuthService         authenticator_service.AuthService
	SimpleAuthService      authenticator_service.AuthService
	EmailOTPAuthService    authenticator_service.AuthService
	MobileOTPAuthService   authenticator_service.AuthService
	MultiFactorAuthService authenticator_service.AuthService
}

func NewAuthResolver(jwtAuthService authenticator_service.AuthService, simpleAuthService authenticator_service.AuthService,
	emailOTPAuthService authenticator_service.AuthService, mobileOTPAuthService authenticator_service.AuthService,
	multiFactorAuthService authenticator_service.AuthService) *AuthResolverImpl {

	if jwtAuthService == nil || simpleAuthService == nil || emailOTPAuthService == nil ||
		mobileOTPAuthService == nil || multiFactorAuthService == nil {
		panic("One or more of the AuthStrategies is/are nil")
	}

	return &AuthResolverImpl{
		JWTAuthService:         jwtAuthService,
		SimpleAuthService:      simpleAuthService,
		EmailOTPAuthService:    emailOTPAuthService,
		MobileOTPAuthService:   mobileOTPAuthService,
		MultiFactorAuthService: multiFactorAuthService,
	}
}

func (a *AuthResolverImpl) AuthStrategyResolver(request authenticator_objects.AuthRequest) authenticator_service.AuthService {
	//if request.IsThirdParty {
	//	return NewThirdPartyAuthStrategy()
	//}

	switch request.AuthType {
	case AuthTypeJWT:
		return a.JWTAuthService
	case AuthTypeSimple:
		return a.SimpleAuthService
	case AuthTypeEmailOTP:
		return a.EmailOTPAuthService
	case AuthTypeMobileOTP:
		return a.MobileOTPAuthService
	case AuthTypeMultiFactor:
		return a.MultiFactorAuthService
	default:
		return a.SimpleAuthService
	}
}
