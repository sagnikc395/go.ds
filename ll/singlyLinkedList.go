package ll

type singlyLinkedList struct {
	head   *Node
	length int
}

func (l *singlyLinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
