package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func pkcs7Pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string to encrypt: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Print("Enter AES key (16, 24 or 32 Bytes): ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	keyLen := len(key)
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		fmt.Printf("Error: Key length must be 16, 24 or 32, but is %d Bytes.\n", keyLen)
		return
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error creating AES cipher: %v\n", err)
		return
	}

	iv := make([]byte, aes.BlockSize)

	paddedText := pkcs7Pad([]byte(text), aes.BlockSize)

	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	fmt.Println("---")
	fmt.Printf("Encrypted (Base64): %s", encoded)
}
