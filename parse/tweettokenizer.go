package parse

import (
	"io"
	"log"
)

type TweetTokenzier struct {
	text               string
	position           int
	currentTokenEntity bool
}

func NewTweetTokenizer(text string) (tokenizer *TweetTokenzier) {
	tokenizer = &TweetTokenzier{text, 0, false}
	return tokenizer
}

func (t *TweetTokenzier) Tokens() (tokens []string) {

	log.Println(t.text)
	for {
		token, err := t.nextToken()
		if err == io.EOF {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func (t *TweetTokenzier) nextToken() (token string, err error) {
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
	return token, nil
}

func (t *TweetTokenzier) isDelimiter(pos int) bool {
	c := t.text[pos]

	switch c {
	case ' ', '\t', '\n':
		t.currentTokenEntity = false
		return true

	case ']', '[', '!', '"', '$', '%', '&', '(', ')', '*', '+', ',', '.', '/',
		';', '<', '>', '=', '?', '\\', '\'', '^', '_', '{', '}', '|', '~', '-', '¬', '·':
		if t.currentTokenEntity == false {
			return true
		}

	case ':':
		if t.currentTokenEntity == false && pos > 3 && t.text[pos-4:pos] != "http" {
			return true
		} else {
			t.currentTokenEntity = true
		}

	case '#':

	}

	return false
}
