package index

type Document struct {
	text   string
	tokens map[string]int
}

func (doc *Document) Text() (text string) {
	return doc.text
}

func (doc *Document) TokenFrequency() (tokenFrequency map[string]int) {
	return doc.tokens
}
