package security

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
)

// HashSHA256 melakukan hashing menggunakan SHA-256
func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// HashArgon2 melakukan hashing dengan Argon2 (lebih aman untuk password)
func HashArgon2(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return hex.EncodeToString(hash)
}
