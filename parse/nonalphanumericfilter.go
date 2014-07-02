package parse

import (
	"log"
	"regexp"
)

var (
	reg *regexp.Regexp
)

type NonAlphanumericFilter struct {
}

func NewNonAlphanumericFilter() *NonAlphanumericFilter {
	var err error
	reg, err = regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return &NonAlphanumericFilter{}
}

func (filter *NonAlphanumericFilter) Filter(input []string) (output []string) {
	for i := range input {
		output = append(output, reg.ReplaceAllString(input[i], ""))
	}
	return output
}
