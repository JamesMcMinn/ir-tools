package parse

import (
	"strings"
)

type LowercaseFilter struct {
}

func NewLowercaseFilter() *LowercaseFilter {
	return &LowercaseFilter{}
}

func (filter *LowercaseFilter) Filter(input []string) (output []string) {
	for i := range input {
		output = append(output, strings.ToLower(input[i]))
	}
	return output
}
