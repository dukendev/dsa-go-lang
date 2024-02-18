package linkedlist

type Data interface {
	int | float32 | int16 | int64 | float64 | string
}

type Node[T Data] struct {
	data T
	next *Node[T]
}

func CreateNode[T Data](value T) *Node[T] {
	return &Node[T]{
		data: value,
		next: nil,
	}
}
