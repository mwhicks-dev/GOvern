package security

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func Encrypt(s string, key []byte) (string, error) {
	/* https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/ */

	// Convert plaintext to bytes
	plaintext := []byte(s)

	// Create cipher from key
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	// Encrypt plaintext
	nonce := make([]byte, gcm.NonceSize())
	cipherbytes := gcm.Seal(nonce, nonce, plaintext, nil)
	ciphertext := string(cipherbytes)

	// Return encrypted plaintext
	return ciphertext, nil
}

func Decrypt(s string, key []byte) (string, error) {
	/* https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/ */

	// Convert ciphertext to bytes
	ciphertext := []byte(s)

	// Create cipher from key
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	// Decrypt ciphertext
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext length less than nonce size")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plainbytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	plaintext := string(plainbytes)

	// Return decrypted ciphertext
	return plaintext, nil
}
