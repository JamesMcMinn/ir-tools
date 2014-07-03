// A tokenizer takes a s string a produces a set of tokens, indivudual
// terms which are split at delimiters.
package parse

type Tokenizer interface {
	Tokens() []string
}
