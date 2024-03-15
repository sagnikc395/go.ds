//we think of it is an tree, but actually an array
// especially good for implementing priority queues

package heap

//a datastructure that can be represented as a complete tree.

//max heap used to fetch the largest key -> the largest key is in the root
// opposite of that is min heap

//actually an area that operate as an tree.

// slice that holds the array
type MaxHeap struct {
	array []int
}

// insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	//maintian heap property
	h.maxHeapifyUp(len(h.array) - 1)
}

// heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		//if parent smaller then swap
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	//index -1 /2 , left child is a odd number and right child is a even number
	return (i - 1) / 2
}

// left child index
func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
