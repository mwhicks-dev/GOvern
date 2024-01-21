package security

import (
	"crypto/sha1"
	"golang.org/x/crypto/pbkdf2"
)

func PasswordToKey(password string, salt string) []byte {
	// Convert salt to bytes
	return pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
}
