package quotes

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quotes struct {
	quotes []string
}

func New(quotesPath string) (*Quotes, error) {
	content, err := os.ReadFile(quotesPath)
	if err != nil {
		return nil, err
	}
	quotes := strings.Split(string(content), "\n")

	return &Quotes{quotes: quotes}, nil
}

func (q *Quotes) GetRandomQuote() (string, error) {
	if len(q.quotes) == 0 {
		return "", fmt.Errorf("no quotes available")
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return q.quotes[rand.Intn(len(q.quotes))], nil
}
