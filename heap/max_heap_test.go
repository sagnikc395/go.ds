package heap

import "testing"

func TestMaxHeapInsert(t *testing.T) {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	if len(m.array) != 3 {
		t.Errorf("expected length=%v,got =%v", 3, len(m.array))

	}
}

func TestMaxHeapExtract(t *testing.T) {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
	}

	if len(m.array) != 5 {
		t.Errorf("expected length=%v,got =%v", 5, len(m.array))

	}
}
 