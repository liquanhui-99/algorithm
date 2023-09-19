package algorithm

// MaxHeap 大顶堆的实现细节
type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: []int{},
	}
}

func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && h.data[index] > h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MaxHeap) ExtractMax() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	mx := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.headifyDown(0)
	return mx, true
}

func (h *MaxHeap) headifyDown(index int) {
	last := len(h.data) - 1
	for {
		left, right := 2*index+1, 2*index+2
		maxIndex := index
		if left <= last && h.data[left] > h.data[maxIndex] {
			maxIndex = left
		}

		if right <= last && h.data[right] > h.data[maxIndex] {
			maxIndex = right
		}

		if index == maxIndex {
			break
		}

		h.data[index], h.data[maxIndex] = h.data[maxIndex], h.data[index]
		maxIndex = index
	}
}
