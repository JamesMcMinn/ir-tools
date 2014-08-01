// LowercaseFilter is badly named - rather than filter out lowercase characters,
// as the name would imply, it converts text to lowercase.
package transform

import (
	"strings"
)

type LowercaseTransform struct {
}

func NewLowercaseTransform() *LowercaseTransform {
	return &LowercaseTransform{}
}

func (filter *LowercaseTransform) ApplyAll(input []string) (output []string) {
	for _, v := range input {
		output = append(output, filter.Apply(v))
	}
	return output
}

func (filter *LowercaseTransform) Apply(input string) (output string) {
	return strings.ToLower(input)
}
