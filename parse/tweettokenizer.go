// TweetTokenizer is a tokenizer designed explicitly for parsing Tweets and other
// Twitter content.
package parse

import (
	"io"
	"log"
	"regexp"
)

type TweetTokenzier struct {
	text               string
	position           int
	currentTokenEntity bool
	CharFilter         *regexp.Regexp
}

func NewTweetTokenizer(text string) (tokenizer *TweetTokenzier) {
	var err error
	reg, err = regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	tokenizer = &TweetTokenzier{text, 0, false, reg}
	return tokenizer
}

func (t *TweetTokenzier) Tokens() (tokens []string) {
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
	if t.currentTokenEntity == false {
		token = reg.ReplaceAllString(token, "")
	}

	t.currentTokenEntity = false

	if token == "" && t.position >= len(t.text)-1 {
		return token, io.EOF
	}
	return token, nil
}

func (t *TweetTokenzier) isDelimiter(pos int) bool {
	c := t.text[pos]

	switch c {
	case ' ', '\t', '\n':
		return true

	case ']', '[', '!', '"', '$', '%', '&', '(', ')', '*', '+', ',', '.', '/',
		';', '<', '>', '=', '?', '\\', '^', '_', '{', '}', '|', '~', '-', '¬', '·':
		if t.currentTokenEntity == false {
			return true
		}

	case ':':
		if t.currentTokenEntity == false && pos > 3 && t.text[pos-4:pos] != "http" {
			return true
		} else {
			t.currentTokenEntity = true
		}

	case '#', '@':
		if t.currentTokenEntity == false && pos == t.position {
			t.currentTokenEntity = true
		} else {
			return true
		}
	}

	return false
}
