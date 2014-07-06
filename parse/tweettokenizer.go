// TweetTokenizer is a tokenizer designed explicitly for parsing Tweets and other
// Twitter content.
package parse

import (
	"io"
)

type TweetTokenzier struct {
	text               string
	position           int
	currentTokenEntity bool
	currentTokenURL    bool
	tokens             []string
	entities           []string
	nonEntities        []string
}

func NewTweetTokenizer(text string) (tokenizer *TweetTokenzier) {
	tokenizer = &TweetTokenzier{text, 0, false, false, []string{}, []string{}, []string{}}
	return tokenizer
}

func (t *TweetTokenzier) Tokens() (tokens []string) {
	if len(t.tokens) == 0 {
		for {
			token, entity, err := t.nextToken()
			if err == io.EOF {
				break
			}
			t.tokens = append(t.tokens, token)
			if entity {
				t.entities = append(t.entities, token)
			} else {
				t.nonEntities = append(t.nonEntities, token)
			}
		}
	}
	return t.tokens
}

func (t *TweetTokenzier) Entities() (entities []string) {
	t.Tokens() // for the side effects only

	return t.entities
}

func (t *TweetTokenzier) NonEntities() (nonentities []string) {
	t.Tokens() // side effects

	return t.nonEntities
}

func (t *TweetTokenzier) AllTokenTypes() (allTokens []string, entities []string, nonEntities []string) {
	return t.Tokens(), t.entities, t.nonEntities
}

func (t *TweetTokenzier) nextToken() (token string, entity bool, err error) {
	t.currentTokenEntity = false
	t.currentTokenURL = false
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
		return token, false, io.EOF
	}

	return token, t.currentTokenEntity, nil
}

func (t *TweetTokenzier) isDelimiter(pos int) bool {
	c := t.text[pos]

	switch c {
	case ' ', '\t', '\n':
		return true

	case ']', '[', '!', '"', '$', '%', '&', '(', ')', '*', '+', ',', '.', '/',
		';', '<', '>', '=', '?', '\\', '^', '_', '{', '}', '|', '~', '-', '¬', '·':
		return t.currentTokenURL != true

	case ':':
		if t.currentTokenEntity == false && pos > 3 && t.text[pos-4:pos] != "http" {
			return true
		} else {
			t.currentTokenEntity = true
			t.currentTokenURL = true
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
