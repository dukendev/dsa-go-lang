package linkedlist

import "fmt"

func (l LinkedList[T]) Search(item T) bool {
	current := l.head
	for current.next != nil {
		if current.data == item {
			fmt.Printf("found %+v", item)
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList[T]) AddAtStart(item T) {
	node := CreateNode[T](item)
	node.next = l.head
	l.size++
	l.head = node
}

func (l *LinkedList[T]) RemoveAtStart() {
	l.head = l.head.next
}

//todo : add list At start

//todo : add list At end
