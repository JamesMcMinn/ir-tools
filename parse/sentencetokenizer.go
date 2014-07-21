// TweetTokenizer is a tokenizer designed explicitly for parsing Tweets and other
// Twitter content.
package parse

import (
	"io"
	"strings"
)

type SentenceTokenizer struct {
	text     string
	tokens   []string
	position int
	comma    bool
}

func NewSentenceTokenizer(text string) (tokenizer *SentenceTokenizer) {
	text = strings.Replace(text, "\n", " ", -1)
	tokenizer = &SentenceTokenizer{
		text:     text,
		tokens:   []string{},
		position: 0,
		comma:    false,
	}
	return tokenizer
}

func (t *SentenceTokenizer) TokenizeOnComma(comma bool) {
	t.comma = comma
}

func (t *SentenceTokenizer) Tokens() (tokens []string) {
	if len(t.tokens) == 0 {
		for {
			token, err := t.nextToken()
			if err == io.EOF {
				break
			}
			t.tokens = append(t.tokens, token)
		}
	}
	return t.tokens
}

func (t *SentenceTokenizer) Longest() (longest string, length int) {
	for _, s := range t.Tokens() {
		if l := len(s); l > length {
			length = l
			longest = s
		}
	}
	return longest, length
}

func (t *SentenceTokenizer) nextToken() (token string, err error) {
	for i := t.position; i < len(t.text); i++ {
		// no new token if the first char after the last delimiter is another delimiter
		if delimiter := t.isDelimiter(i); delimiter && t.position == i {
			t.position++
		} else if delimiter {
			token = t.text[t.position:i]
			t.position = i + 1
			break
		} else if i == (len(t.text) - 1) { // We're at the last char and it's not a delimiter, so we must include it
			token = t.text[t.position : i+1]
			t.position = i + 1
			break
		}
	}

	if token == "" && t.position >= len(t.text)-1 {
		return token, io.EOF
	}

	token = strings.Trim(token, " ")
	return token, nil
}

func (t *SentenceTokenizer) isDelimiter(pos int) bool {
	c := t.text[pos]

	switch c {
	case '!', '?', '|', '·':
		return true

	case '-', '~', '.':
		return (pos+1) == len(t.text) || t.text[pos+1] == ' '

	// case ':':
	// 	return (pos+1) == len(t.text) || t.text[pos+1] != '/'

	case ',':
		return t.comma && ((pos+1) == len(t.text) || t.text[pos+1] == ' ')
	}

	return false
}
