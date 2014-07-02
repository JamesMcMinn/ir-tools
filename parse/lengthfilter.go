package parse

type LengthFilter struct {
	minLength int
}

func NewLengthFilter(minLength int) *LengthFilter {
	return &LengthFilter{minLength}
}

func (filter *LengthFilter) Filter(input []string) (output []string) {
	for i := range input {
		if len(input[i]) < filter.minLength {
			continue
		}

		output = append(output, input[i])
	}
	return output
}
