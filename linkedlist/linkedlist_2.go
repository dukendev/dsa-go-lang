package linkedlist

func (l *LinkedList[T]) AppendAtStart(list *LinkedList[T]) {
	if l == nil || list == nil || list.head == nil {
		return
	}
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = l.head
	l.head = list.head
}

// todo : add list At end
func (l *LinkedList[T]) AppendAtEnd(list *LinkedList[T]) {
	if list == nil || list.head == nil {
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = list.head
}

//todo : reverse list

//todo : sort list

//todo : merge two sorted list

//todo : detech cycle

//todo : find cycle

//todo : reverse sublist
