package filter

import (
	"github.com/jamesmcminn/ir-tools"
	"log"
)

type StopWordFilter struct {
	stopwords map[string]bool
}

func NewStopWordFilter(stopwordFile string) *StopWordFilter {
	terms, err := irtools.ReadLines(stopwordFile)

	if err != nil {
		log.Fatal(err)
	}
	stopwords := map[string]bool{}

	for _, term := range terms {
		stopwords[term] = true
	}

	return &StopWordFilter{stopwords}
}

func (filter *StopWordFilter) Filter(input []string) (output []string) {
	for i := range input {
		if _, found := filter.stopwords[input[i]]; found == true {
			continue
		}

		output = append(output, input[i])
	}
	return output
}
