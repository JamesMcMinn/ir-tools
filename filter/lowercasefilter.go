// LowercaseFilter is badly named - rather than filter out lowercase characters,
// as the name would imply, it converts text to lowercase.
package filter

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
