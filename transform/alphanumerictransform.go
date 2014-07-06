package transform

import (
	"log"
	"regexp"
)

var (
	reg *regexp.Regexp
)

type AlphanumericTransform struct {
}

func NewAlphanumericTransform() *AlphanumericTransform {
	var err error
	reg, err = regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return &AlphanumericTransform{}
}

func (filter *AlphanumericTransform) Apply(input []string) (output []string) {
	for i := range input {
		output = append(output, reg.ReplaceAllString(input[i], ""))
	}
	return output
}
