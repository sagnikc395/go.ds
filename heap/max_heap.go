//we think of it is an tree, but actually an array
// especially good for implementing priority queues

package heap

import "fmt"

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

// extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array)

	//edge case for empty array
	if len(h.array) == 0 {
		fmt.Printf("Cannot extract because array length is 0\n")
		return -1
	}

	h.array[0] = h.array[l-1]
	//make the slice shorter to get the rid of the last one
	h.array = h.array[:l]

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	//heapify from top to bottom
	//loop while index has at least one child
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		//when left is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			//when left child is larger
			childToCompare = l
		} else {
			//when right child is larger
			childToCompare = r
		}
		//compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			//stop the heapify process and return
			return
		}
	}
}
