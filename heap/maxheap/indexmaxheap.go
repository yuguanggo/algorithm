package maxheap

type IndexMaxHeap struct {
	data []int
	indexes []int
	count int
	capacity int
}

func Init(n int) *IndexMaxHeap {
	return &IndexMaxHeap{
		data:make([]int,n),
		indexes:make([]int,n),
		count:0,
		capacity:n,
	}
}

func (h *IndexMaxHeap) IsEmpty() bool  {
	return h.count==0
}
// 向最大索引堆中插入一个新的元素, 新元素的索引为index, v
// 传入的index对用户而言,是从0索引的
func (h *IndexMaxHeap) Insert(index int,v int)  {
	h.count++
	h.data[index] = v
	h.indexes = append(h.indexes,index)
	h.shiftUp(index)
}

func (h *IndexMaxHeap) shiftUp(index int)  {
	
}


