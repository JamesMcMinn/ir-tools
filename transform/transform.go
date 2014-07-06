package transform

type transform interface {
	Apply(input []string) (output []string)
}
