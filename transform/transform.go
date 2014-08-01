package transform

type Transform interface {
	ApplyAll(input []string) (output []string)
	Apply(input string) (output string)
}
