package three_in_one

/*
	3.1: Three in One: Describe how you could use a single array to implement three stacks.
*/

type ThreeStack[T any] struct {
	List   []T
	Length int
}

func CreateThreeStack[T any](l int) *ThreeStack[T] {
	return &ThreeStack[T]{
		List:   make([]T, l),
		Length: l,
	}
}
