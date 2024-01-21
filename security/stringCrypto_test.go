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
	text, err := encrypt(initial, key)
	if err != nil {
		t.Fatalf("encryption failed")
	}

	text, err = decrypt(text, key)
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
	text1, err := encrypt(initial, key1)
	if err != nil {
		t.Fatalf("encryption 1 failed")
	}

	text2, err := encrypt(initial, key2)
	if err != nil {
		t.Fatalf("encryption 2 failed")
	}

	text3, err := encrypt(initial, key1)
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
