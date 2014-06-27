package parse

type filter interface {
	Filter(input []string) (output []string)
}
