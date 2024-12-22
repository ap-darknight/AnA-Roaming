package symmetric_strategies

import (
	crypto_objects "AnA-Roaming/ana-authenticator/dto-layer/crypto-objects"
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/chacha20"
)

// ChaCha20CryptographyService provides encryption and decryption using ChaCha20
type ChaCha20CryptographyService struct {
	Config *config_dto.Config
}

const (
	ChaCha20KeySize = 32 // ChaCha20 key size: 256 bits
	NonceSize       = 12 // ChaCha20 nonce size: 96 bits
)

func NewChaCha20CryptographyService(config *config_dto.Config) *ChaCha20CryptographyService {
	return &ChaCha20CryptographyService{Config: config}
}

// EncryptDataWithKey encrypts the given CryptoObject using ChaCha20 with a generated key
func (c ChaCha20CryptographyService) EncryptDataWithKey(data crypto_objects.CryptoObject) (string, error) {

	key := c.Config.Crypto.Chacha20.EncryptionKey

	// Validate key size
	if len(key) != ChaCha20KeySize {
		return "", errors.New("key length must be 32 bytes for Encryption")
	}

	// Convert the CryptoObject to bytes
	plainData, err := data.ToBytes()
	if err != nil {
		return "", err
	}

	// Generate a nonce
	nonce := make([]byte, NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// Create the ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher([]byte(key), nonce)
	if err != nil {
		return "", err
	}

	// Encrypt the data
	ciphertext := make([]byte, len(plainData))
	cipher.XORKeyStream(ciphertext, plainData)

	// Prepend the nonce to the ciphertext
	ciphertextWithNonce := append(nonce, ciphertext...)

	// Return the base64-encoded ciphertext
	return base64.URLEncoding.EncodeToString(ciphertextWithNonce), nil
}

// DecryptDataWithKey decrypts the given encrypted string using ChaCha20 with a provided key
func (c ChaCha20CryptographyService) DecryptDataWithKey(data string) (interface{}, error) {

	key := c.Config.Crypto.Chacha20.EncryptionKey

	// Validate key size
	if len(key) != ChaCha20KeySize {
		return nil, errors.New("key length must be 32 bytes for Decryption")
	}

	// Decode the base64-encoded ciphertext
	ciphertextWithNonce, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// Extract the nonce and ciphertext
	if len(ciphertextWithNonce) < NonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertextWithNonce[:NonceSize], ciphertextWithNonce[NonceSize:]

	// Create the ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher([]byte(key), nonce)
	if err != nil {
		return nil, err
	}

	// Decrypt the data
	plainData := make([]byte, len(ciphertext))
	cipher.XORKeyStream(plainData, ciphertext)

	// Convert the decrypted bytes back to a CryptoObject
	obj, err := crypto_objects.FromBytes(plainData)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
