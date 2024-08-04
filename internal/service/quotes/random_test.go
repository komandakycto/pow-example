package quotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuotes_GetRandomQuote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		quotes  []string
		wantErr bool
	}{
		{
			name:    "No quotes available",
			quotes:  []string{},
			wantErr: true,
		},
		{
			name:    "Single quote available",
			quotes:  []string{"Guard well your thoughts when alone and your words when accompanied."},
			wantErr: false,
		},
		{
			name:    "Multiple quotes available",
			quotes:  []string{"Guard well your thoughts when alone and your words when accompanied.", "I like to listen. I have learned a great deal from listening carefully. Most people never listen."},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q := &Quotes{
				quotes: tt.quotes,
			}
			got, err := q.GetRandomQuote()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Contains(t, tt.quotes, got)
			}
		})
	}
}
