package security

import (
	"strings"
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	// Sample data
	initial := "The quick brown fox jumps over the lazy dog, and a cat watches."
	key := []byte("TEST KEY 0000001")

	// Encrypted and decrypt sample data
	text, err := Encrypt(initial, key)
	if err != nil {
		t.Fatalf("encryption failed")
	}

	text, err = Decrypt(text, key)
	if err != nil {
		t.Fatalf("decryption failed")
	}

	// Ensure correct behavior
	if strings.Compare(initial, text) != 0 {
		t.Fatalf("final and initial strings differ")
	}
}

func TestEncryption(t *testing.T) {
	// Sample data
	initial := "The quick brown fox jumps over the lazy dog, and a cat watches."
	key1 := []byte("TEST KEY 0000001")
	key2 := []byte("TEST KEY 0000002")

	// Encrypt sample data
	text1, err := Encrypt(initial, key1)
	if err != nil {
		t.Fatalf("encryption 1 failed")
	}

	text2, err := Encrypt(initial, key2)
	if err != nil {
		t.Fatalf("encryption 2 failed")
	}

	text3, err := Encrypt(initial, key1)
	if err != nil {
		t.Fatalf("encryption 3 failed")
	}

	// Ensure correct behavior
	if strings.Compare(text1, text2) == 0 {
		t.Fatalf("separate keys encrypt to same string")
	}
	if strings.Compare(text1, text3) != 0 {
		t.Fatalf("same key encrypts to separate strings")
	}
}

func TestDecryption(t *testing.T) {
	// Sample data
	initial := "The quick brown fox jumps over the lazy dog, and a cat watches."
	key1 := []byte("TEST KEY 0000001")
	key2 := []byte("TEST KEY 0000002")

	// Encrypt sample data
	text1, _ := Encrypt(initial, key1)
	text2, _ := Encrypt(initial, key2)

	// Decrypt sample data
	text1, _ = Decrypt(initial, key2)
	text2, _ = Decrypt(initial, key1)

	// Ensure correct behavior
	if strings.Compare(initial, text1) == 0 || strings.Compare(initial, text2) == 0 {
		t.Fatalf("different key decrypts encrypted string")
	}
}

func TestAesKdf(t *testing.T) {
	// Sample data
	key := PasswordToKey("password123", "master")
	keySame := PasswordToKey("password123", "master")
	keyDiff := PasswordToKey("password321", "master")
	keyDiffSalt := PasswordToKey("password123", "pwencrypt")

	// Ensure correct behavior
	if strings.Compare(string(key), string(keySame)) != 0 {
		t.Fatalf("same password and salt hash to different key")
	}
	if strings.Compare(string(key), string(keyDiffSalt)) == 0 {
		t.Fatalf("same password and different salt hash to same key")
	}
	if strings.Compare(string(key), string(keyDiff)) == 0 {
		t.Fatalf("different password and same key hash to same key")
	}
}
