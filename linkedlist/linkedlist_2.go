package linkedlist

// todo : add list At start
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

//todo : add list At end

//todo : reverse list

//todo : sort list

//todo : merge two sorted list

//todo : detech cycle

//todo : find cycle

//todo : reverse sublist
