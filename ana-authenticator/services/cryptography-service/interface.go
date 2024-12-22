package cryptography_service

import crypto_objects "AnA-Roaming/ana-authenticator/dto-layer/crypto-objects"

// CryptographyService is the interface for the cryptography service
type CryptographyService interface {
	EncryptDataWithKey(data crypto_objects.CryptoObject) (string, error)
	DecryptDataWithKey(data string) (interface{}, error)
}
