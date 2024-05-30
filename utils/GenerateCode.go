package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

func GenerateChallengeCode() string {
	// Generate a 128-character code verifier
	possible := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
	verifier := make([]byte, 128)
	for i := range verifier {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(possible))))
		verifier[i] = possible[randomIndex.Int64()]
	}

	// Compute the SHA-256 hash of the verifier
	sha256Hash := sha256.Sum256(verifier)

	// Base64 URL encode the hash
	challenge := base64.RawURLEncoding.EncodeToString(sha256Hash[:])

	return challenge
}
