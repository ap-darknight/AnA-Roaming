package symmetric_strategies

import (
	crypto_objects "AnA-Roaming/ana-authenticator/dto-layer/crypto-objects"
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AESCryptographyService provides AES encryption and decryption methods
type AESCryptographyService struct {
	Config *config_dto.Config
}

const (
	AESKeySize = 32 // ChaCha20 key size: 256 bits
)

func NewAESCryptographyService(config *config_dto.Config) *AESCryptographyService {
	return &AESCryptographyService{Config: config}
}

// EncryptDataWithKey encrypts the given data using AES encryption with the provided key
func (a AESCryptographyService) EncryptDataWithKey(data crypto_objects.CryptoObject) (string, error) {

	// Convert the key to a byte slice
	keyBytes := []byte(a.Config.Crypto.Aes256.EncryptionKey)
	if len(keyBytes) != AESKeySize {
		return "", errors.New("key length must be 32 bytes for Encryption")
	}

	// Convert the data to a byte slice
	plainData, err := data.ToBytes()
	if err != nil {
		return "", err
	}

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Use GCM (Galois/Counter Mode) for encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data and prepend the nonce to the ciphertext
	ciphertext := aesGCM.Seal(nonce, nonce, plainData, nil)

	// Return the base64-encoded ciphertext
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptDataWithKey decrypts the given data using AES decryption with the provided key
func (a AESCryptographyService) DecryptDataWithKey(data string) (interface{}, error) {
	// Convert the key to a byte slice
	keyBytes := []byte(a.Config.Crypto.Aes256.EncryptionKey)
	if len(keyBytes) != AESKeySize {
		return nil, errors.New("key length must be 32 bytes for Decryption")
	}

	// Decode the base64-encoded ciphertext
	ciphertext, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	// Use GCM (Galois/Counter Mode) for decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Extract the nonce and the actual ciphertext
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plainData, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice back to the original format
	return crypto_objects.FromBytes(plainData) // Assuming a FromBytes method exists in CryptoObject
}
