package transform

import (
	"log"
	"regexp"
)

type AlphanumericTransform struct {
	reg *regexp.Regexp
}

// Create a new Alphanumeric Transformer, which removes any non-alhpanumeric
// chracters.
func NewAlphanumericTransform(extra string) *AlphanumericTransform {
	var err error
	var reg *regexp.Regexp

	reg, err = regexp.Compile("[^A-Za-z0-9" + extra + "]+")
	if err != nil {
		log.Fatal(err)
	}

	return &AlphanumericTransform{reg: reg}
}

func (filter *AlphanumericTransform) ApplyAll(input []string) (output []string) {
	for i := range input {
		output = append(output, filter.Apply(input[i]))
	}
	return output
}

func (filter *AlphanumericTransform) Apply(input string) (output string) {
	return filter.reg.ReplaceAllString(input, "")
}
