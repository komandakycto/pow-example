package hashcash

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashcash_GenerateChallenge(t *testing.T) {
	h := New(1)

	challenge1 := h.GenerateChallenge()
	challenge2 := h.GenerateChallenge()

	assert.NotEmpty(t, challenge1, "Challenge should not be empty")
	assert.NotEmpty(t, challenge2, "Challenge should not be empty")
	assert.NotEqual(t, challenge1, challenge2, "Challenges should be different")
}

func TestHashcash_VerifyPoW(t *testing.T) {
	h := New(1)

	challenge := h.GenerateChallenge()
	nonce := "12345"

	// Create a valid proof of work
	rand.Seed(42) // Setting a fixed seed for deterministic behavior
	hash := sha256.Sum256([]byte(challenge + nonce))
	hashStr := hex.EncodeToString(hash[:])
	for !strings.HasPrefix(hashStr, "0") {
		nonce = strconv.Itoa(rand.Int())
		hash = sha256.Sum256([]byte(challenge + nonce))
		hashStr = hex.EncodeToString(hash[:])
	}

	assert.True(t, h.VerifyPoW(challenge, nonce), "Valid PoW should return true")

	// Test with an invalid proof of work
	assert.False(t, h.VerifyPoW(challenge, "invalid_nonce"), "Invalid PoW should return false")
}
