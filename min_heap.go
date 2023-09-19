package algorithm

// MinHeap 小顶堆的实现细节
type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	// 开始排序
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	parent := (index - 1) / 2
	// 子节点的值比父节点的值大，才能替换
	if index > 0 && h.data[index] < h.data[parent] {
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// ExtractMin 移除小顶堆中的最小值
func (h *MinHeap) ExtractMin() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}
	mn := h.data[0]
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]
	// 开始对下标为0的数据进行重新降层排序
	h.heapipyDown(0)
	return mn, true
}

func (h *MinHeap) heapipyDown(index int) {
	last := len(h.data) - 1
	for {
		left, right := 2*index+1, 2*index+2
		smallIndex := index
		// 左子节点比右子节点小，那么只对左子节点进行排序
		if left <= last && h.data[left] < h.data[smallIndex] {
			smallIndex = left
		}
		// 左子节点比右子节点小，那么只对右子节点进行排序
		if right <= last && h.data[right] < h.data[smallIndex] {
			smallIndex = right
		}

		if smallIndex == index {
			break
		}

		h.data[index], h.data[smallIndex] = h.data[smallIndex], h.data[index]
		index = smallIndex
	}
}
