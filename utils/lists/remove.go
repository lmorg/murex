package lists

import "fmt"

func RemoveOrdered[T any](slice []T, i int) ([]T, error) {
	switch {
	case i >= len(slice):
		fallthrough
	case i < 0:
		return nil, fmt.Errorf("index out of bounds: len=%d, index=%d", len(slice), i)
	default:
		return append(slice[:i], slice[i+1:]...), nil
	}
}

func RemoveUnordered[T any](slice []T, i int) ([]T, error) {
	switch {
	case i >= len(slice):
		fallthrough
	case i < 0:
		return nil, fmt.Errorf("index out of bounds: len=%d, index=%d", len(slice), i)
	default:
		slice[i] = slice[len(slice)-1]
		return slice[:len(slice)-1], nil
	}
}
