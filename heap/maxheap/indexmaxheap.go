package maxheap

type IndexMaxHeap struct {
	data     []int
	indexes  []int
	count    int
}

func Init(n int) *IndexMaxHeap {
	return &IndexMaxHeap{
		data:     make([]int, n),
		indexes:  make([]int, 0),
		count:    0,
	}
}

func (h *IndexMaxHeap) IsEmpty() bool {
	return h.count == 0
}

// 向最大索引堆中插入一个新的元素, 新元素的索引为index, v
// 传入的index对用户而言,是从0索引的
func (h *IndexMaxHeap) Insert(index int, v int) {
	h.count++
	h.data[index] = v
	h.indexes = append(h.indexes, index)
	h.shiftUp(len(h.indexes) - 1)
}

//k是indexes数组中最后一位
func (h *IndexMaxHeap) shiftUp(k int) {
	for k > 0 && h.data[h.indexes[(k-1)/2]] < h.data[h.indexes[k]] {
		h.swap((k-1)/2, k)
		k = (k - 1) / 2
	}
}

//交换indexes数组中 i,j的位置
func (h *IndexMaxHeap) swap(i, j int) {
	h.indexes[i], h.indexes[j] = h.indexes[j], h.indexes[i]
}

func (h *IndexMaxHeap) ExtractMax() int {
	maxIndex := h.indexes[0]
	//首尾交换下位置
	h.swap(0, len(h.indexes)-1)
	//去掉末尾元素
	h.indexes = h.indexes[:len(h.indexes)-1]
	h.count--
	h.shiftDown(0)
	return h.data[maxIndex]
}

func (h *IndexMaxHeap) shiftDown(k int) {
	//左子树
	n := len(h.indexes)
	for 2*k + 1 < n {
		j := 2*k + 1

		if j+1 < n && h.data[h.indexes[j]] < h.data[h.indexes[j+1]] {
			j++
		}
		if h.data[h.indexes[k]] > h.data[h.indexes[j]] {
			break
		}
		h.swap(k, j)
		k = j
	}
}

func (h *IndexMaxHeap) Change(i,v int)  {
	h.data[i] = v
	for j:=0;i<len(h.indexes);j++{
		if h.indexes[j]==i{
			h.shiftUp(j)
			h.shiftDown(j)
			return
		}
	}
}
