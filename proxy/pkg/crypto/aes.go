package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"
)

func DecryptPayload(key []byte, iv []byte, encryptedString string) (string, error) {
	cipherText, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	// prepare decryption based on key and iv
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	stream := cipher.NewCFBDecrypter(block, iv)

	// decrypt
	stream.XORKeyStream(cipherText, cipherText)

	// fix
	result := strings.ReplaceAll(string(cipherText), "\x00", "")
	return result, nil
}
