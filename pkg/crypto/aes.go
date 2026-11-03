package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// PKCS7Unpad removes PKCS#7 padding from decrypted data.
func PKCS7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	padLen := int(data[length-1])
	if padLen == 0 || padLen > length {
		return nil, fmt.Errorf("invalid padding length: %d", padLen)
	}
	for i := length - padLen; i < length; i++ {
		if data[i] != byte(padLen) {
			return nil, fmt.Errorf("invalid padding byte")
		}
	}
	return data[:length-padLen], nil
}

// DecryptAESCBCBase64 decrypts a Base64 encoded ciphertext encrypted with AES-CBC.
// The IV is expected to be prepended in the first 16 bytes of the ciphertext.
func DecryptAESCBCBase64(cipherBase64, keyText string) (string, error) {
	cipherData, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	key := []byte(keyText)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("invalid AES key size: %d bytes. Must be 16, 24, or 32", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	if len(cipherData) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short (minimum %d bytes)", aes.BlockSize)
	}

	iv := cipherData[:aes.BlockSize]
	ciphertext := cipherData[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = PKCS7Unpad(plaintext)
	if err != nil {
		return "", fmt.Errorf("failed to unpad plaintext: %w", err)
	}

	return string(plaintext), nil
}

