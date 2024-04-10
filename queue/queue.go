package queue

import (
	"com.github/dukendev/basics/common"
)

type Queue[T common.Data] struct {
	arr  []T
	size int
}

func (q *Queue[T]) Enqueue(item T) {
	if q.size == 0 {
		q.arr = make([]T, 0)
	}
	q.size++
	q.arr = append(q.arr, item)
}

func (q *Queue[T]) Dequeue() *T {
	if q.size == 0 {
		return nil
	}
	q.size--
	item := q.arr[0]
	q.arr = q.arr[1:]
	return &item
}

func (q Queue[T]) Show() {
	for i := 0; i < q.size; i++ {
		print(q.arr[i])
		print(" ")
	}
	println()
}
