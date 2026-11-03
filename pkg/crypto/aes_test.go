package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"testing"
)

func encryptAESCBC(plaintext, keyText string) (string, error) {
	key := []byte(keyText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	padLen := aes.BlockSize - (len(plaintext) % aes.BlockSize)
	padded := append([]byte(plaintext), make([]byte, padLen)...)
	for i := len(plaintext); i < len(padded); i++ {
		padded[i] = byte(padLen)
	}

	ciphertext := make([]byte, aes.BlockSize+len(padded))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padded)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func TestDecryptAESCBCBase64(t *testing.T) {
	key := "1234567890123456" // 16 bytes
	original := "secret-solace-password-123!"

	encrypted, err := encryptAESCBC(original, key)
	if err != nil {
		t.Fatalf("failed to encrypt: %v", err)
	}

	decrypted, err := DecryptAESCBCBase64(encrypted, key)
	if err != nil {
		t.Fatalf("failed to decrypt: %v", err)
	}

	if decrypted != original {
		t.Errorf("got %q, want %q", decrypted, original)
	}
}

