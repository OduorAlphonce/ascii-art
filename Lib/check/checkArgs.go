package check

func Args(input []string) (string, bool) {
	inputLength := len(input)

	if inputLength != 2 {
		return "Did you pas a word like: go run . \"some-text?\"", false
	}
	data := input[1]
	if len(data) < 1 {
		return "I can't work with an empty string", false
	}
	return data, true
}