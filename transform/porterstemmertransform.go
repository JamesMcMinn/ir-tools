package transform

import (
	"github.com/dchest/stemmer"
	"github.com/dchest/stemmer/porter2"
)

type PorterStemmer struct {
	stemmer stemmer.Stemmer
}

func NewPorterStemmer() *PorterStemmer {
	return &PorterStemmer{porter2.Stemmer}
}

func (transform *PorterStemmer) Apply(input []string) (output []string) {
	for i := range input {
		output = append(output, transform.stemmer.Stem(input[i]))
	}
	return output
}
