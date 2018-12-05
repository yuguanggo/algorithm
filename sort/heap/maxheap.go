package heap

type MaxHeap  []int

func (h MaxHeap) Swap(i,j int)  {
	h[i],h[j] = h[j],h[i]
}

func (h MaxHeap) Size() int  {
	return len(h)
}

func (h MaxHeap) IsEmpty() bool  {
	return len(h)==0
}


