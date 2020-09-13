// Package accumulate implements the accumulate operation.
package accumulate

// Function Accumulate, given a collection and an operation to perform one ach element of the
// collection, returns a new collection containing the result of applying that operation to each
// element of the input collection.
func Accumulate(inputs []string, operation func(string) string) []string {
	outputs := make([]string, 0, len(inputs))
	for _, input := range inputs {
		outputs = append(outputs, operation(input))
	}
	return outputs
}
