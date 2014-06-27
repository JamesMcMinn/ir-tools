package parse

import (
	"strings"
)

type WhitespaceTokenier struct {
	text string
}

func NewWhitespaceTokenier(text string) (tokenizer *WhitespaceTokenier) {
	tokenizer = new(WhitespaceTokenier)
	tokenizer.text = text
	return tokenizer
}

func (t *WhitespaceTokenier) Tokens() []string {
	return strings.Fields(t.text)
}
