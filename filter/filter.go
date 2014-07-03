package filter

type filter interface {
	Filter(input []string) (output []string)
}
