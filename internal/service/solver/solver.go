// Package solver provides a simple proof-of-work solver.
package solver

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

type Solver struct {
	difficulty int
}

func New(difficulty int) *Solver {
	return &Solver{difficulty: difficulty}
}

func (s *Solver) SolvePoW(challenge string) string {
	prefix := strings.Repeat("0", s.difficulty)
	nonce := 0
	for {
		hash := sha256.Sum256([]byte(challenge + strconv.Itoa(nonce)))
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, prefix) {
			return strconv.Itoa(nonce)
		}
		nonce++
	}
}
