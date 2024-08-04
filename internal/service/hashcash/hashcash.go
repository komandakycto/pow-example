// Package hashcash provides a simple proof-of-work implementation.
package hashcash

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Hashcash struct {
	difficulty int
}

func New(difficulty int) *Hashcash {
	return &Hashcash{difficulty: difficulty}
}

func (h *Hashcash) GenerateChallenge() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return strconv.Itoa(rand.Int())
}

func (h *Hashcash) VerifyPoW(challenge, nonce string) bool {
	hash := sha256.Sum256([]byte(challenge + nonce))
	hashStr := hex.EncodeToString(hash[:])

	return strings.HasPrefix(hashStr, strings.Repeat("0", h.difficulty))
}
