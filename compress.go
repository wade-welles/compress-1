package compress

import (
	"io"
)

var algorithms = make(map[string]Algorithm)

// Algorithm interface.
type Algorithm interface {
	NewEncoder(writer io.Writer) (Encoder, error)
	NewDecoder(reader io.Reader) (Decoder, error)
}

// Register algorithm.
func Register(name string, algorithm Algorithm) {
	algorithms[name] = algorithm
}

// Registered algorithm.
func Registered(name string) bool {
	_, ok := algorithms[name]
	if !ok {
		return false
	}
	return true
}

// Algorithms registered.
func Algorithms() []string {
	l := []string{}
	for k := range algorithms {
		l = append(l, k)
	}
	return l
}