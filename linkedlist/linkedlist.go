package linkedlist

import (
	"com.github/dukendev/basics/common"
	"fmt"
)

type LinkedList[T common.Data] struct {
	head *Node[T]
	size int
}

func (l *LinkedList[T]) Add(node *Node[T]) {
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current != nil && current.next != nil {
			current = current.next
		}
		current.next = node
	}
	l.size++
}

func (l LinkedList[T]) Show() {
	current := l.head
	for current != nil {
		fmt.Printf("%+v -> ", current.data)
		if current.next == nil {
			fmt.Print("EOL")
		}
		current = current.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) Remove(item T) {
	if l.head.data == item {
		l.head = l.head.next
	} else {
		current := l.head
		for current.next.data != item && current.next.next != nil {
			current = current.next
		}
		if current.next.next == nil {
			if current.next.data == item {
				current.next = nil
			}
			return
		}
		current.next = current.next.next
	}
}
