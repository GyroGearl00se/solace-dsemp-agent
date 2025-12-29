package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func pkcs7Pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	padLen := int(data[length-1])
	if padLen > length {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:length-padLen], nil
}

func encryptString(text string, key string) (string, error) {
	keyLen := len(key)
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return "", fmt.Errorf("key length must be 16, 24 or 32, but is %d Bytes", keyLen)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("creating AES cipher: %w", err)
	}

	paddedText := pkcs7Pad([]byte(text), aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(paddedText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("creating IV: %w", err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedText)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

func decryptString(cipherBase64, keyText string) (string, error) {
	// Decode Base64
	cipherData, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return "", err
	}

	// Key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
	key := []byte(keyText)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("invalid AES key size: %d bytes. Must be 16, 24, or 32", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("creating AES cipher: %w", err)
	}

	if len(cipherData) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// Extract IV from the beginning of the ciphertext
	iv := cipherData[:aes.BlockSize]
	ciphertext := cipherData[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func encrypt() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string to encrypt: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Print("Enter AES key (16, 24 or 32 Bytes): ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	encrypted, err := encryptString(text, key)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("---")
	fmt.Printf("Encrypted (Base64): %s\n", encrypted)
}

func decrypt() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Base64 string to decrypt: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Print("Enter AES key (16, 24 or 32 Bytes): ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	decrypted, err := decryptString(text, key)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("---")
	fmt.Printf("Decrypted: %s\n", decrypted)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [encrypt|decrypt]")
		return
	}

	command := os.Args[1]

	switch command {
	case "encrypt":
		encrypt()
	case "decrypt":
		decrypt()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Usage: go run main.go [encrypt|decrypt]")
	}
}
