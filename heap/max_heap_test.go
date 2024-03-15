package heap

import "testing"

func TestMaxHeap(t *testing.T) {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	if len(m.array) != 3 {
		t.Errorf("expected length=%v,got %v", 3, len(m.array))
		
	}

}
