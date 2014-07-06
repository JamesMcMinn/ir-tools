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

func (filter *LowercaseTransform) Apply(input []string) (output []string) {
	for i := range input {
		output = append(output, strings.ToLower(input[i]))
	}
	return output
}
