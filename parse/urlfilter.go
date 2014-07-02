// URL Filter removes URLs from the inpout tokens, and returns the remaining tokens
package parse

import (
	"strings"
)

type URLFilter struct {
}

func NewURLFilter() *URLFilter {
	return &URLFilter{}
}

func (filter *URLFilter) Filter(input []string) (output []string) {
	for i := range input {
		if strings.HasPrefix(strings.ToLower(input[i]), "http://") {
			continue
		}

		output = append(output, input[i])
	}
	return output
}
