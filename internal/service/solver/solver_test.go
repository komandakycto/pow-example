package solver

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePoW(t *testing.T) {
	difficulty := 2
	solver := New(difficulty)
	challenge := "test-challenge"

	nonce := solver.SolvePoW(challenge)
	hash := sha256.Sum256([]byte(challenge + nonce))
	hashStr := hex.EncodeToString(hash[:])

	assert.True(t, strings.HasPrefix(hashStr, strings.Repeat("0", difficulty)), "Hash should have the required number of leading zeros")
}

func TestSolvePoWWithDifferentDifficulties(t *testing.T) {
	challenges := []string{"challenge1", "challenge2", "challenge3"}
	difficulties := []int{1, 2, 3, 4}

	for _, challenge := range challenges {
		for _, difficulty := range difficulties {
			solver := New(difficulty)
			nonce := solver.SolvePoW(challenge)
			hash := sha256.Sum256([]byte(challenge + nonce))
			hashStr := hex.EncodeToString(hash[:])

			assert.True(t, strings.HasPrefix(hashStr, strings.Repeat("0", difficulty)), "Hash should have the required number of leading zeros")
		}
	}
}
