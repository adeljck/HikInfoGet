package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func DecryptData(base string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return "", err
	}

	// Check magic number
	if decode[0] != 17 {
		return "", fmt.Errorf("Invalid magic number")
	}

	// Extract size and salt
	size := int(decode[1])<<8 | int(decode[2])
	salt := decode[4 : 4+size]

	// Extract IV and data
	iv := decode[20 : 20+16]
	data := decode[36 : 36+size]

	// Derive AES key from password and salt using PBKDF2
	key := pbkdf2.Key([]byte("Abc123@&$++Hik45"), salt, 10000, 32, sha256.New)

	// Create AES cipher in CBC mode with PKCS5Padding
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt data
	mode.CryptBlocks(data, data)

	// Remove padding
	unpad := func(src []byte) []byte {
		length := len(src)
		unpadding := int(src[length-1])
		return src[:(length - unpadding)]
	}

	decryptedData := unpad(data)

	// Convert decrypted data to UTF-8 string
	return string(decryptedData), nil
}
