package ll

import "testing"

func TestSinglyLinkedlist(t *testing.T) {
	mylist := &singlyLinkedList{}
	node1 := &Node{data: 1}

	mylist.prepend(node1)

	if mylist.head != node1 {
		t.Errorf("Expected head to be %v, got %v", node1, mylist.head)
	}
	if mylist.length != 1 {
		t.Errorf("Expected length to be 1, got %v", mylist.length)
	}
}
