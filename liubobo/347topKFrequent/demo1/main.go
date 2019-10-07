package main


func topKFrequent(nums []int, k int) []int {
	record:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		record[nums[i]]++
	}
	h:=MinHeap{}
	for key,val:=range record{
		if h.Size()>=k{
			min:=h.GetMin()
			if min<val{
				h.ExtractMin()
				h.Insert(node{key,val})
			}
		}else{
			h.Insert(node{key,val})
		}
	}
	res:=make([]int,k)
	for i:=k-1;i>=0;i--{
		res[i]=h.ExtractMin()
	}
	return res
}

type node struct {
	val int
	count int
}

type MinHeap []node

func (h *MinHeap)Size() int  {
	return len(*h)
}
func (h *MinHeap)Insert(item node)  {
	*h=append(*h,item)
	h.shiftUp(h.Size()-1)
}

func (h *MinHeap)GetMin() int  {
	return (*h)[0].count
}

func (h *MinHeap)ExtractMin() int  {
	item:=(*h)[0]
	h.swap(0,len(*h)-1)
	*h=(*h)[:len(*h)-1]
	h.shiftDown(0)
	return item.val
}


func (h *MinHeap)swap(i,j int)  {
	(*h)[i],(*h)[j]=(*h)[j],(*h)[i]
}

func (h *MinHeap)shiftUp(k int)  {
	for k>0&&(*h)[k].count<(*h)[(k-1)/2].count{
		h.swap(k,(k-1)/2)
		k=(k-1)/2
	}
}

func (h *MinHeap)shiftDown(k int)  {
	l:=2*k+1
	changK:=l
	if l+1<h.Size()&&(*h)[l].count>(*h)[l+1].count{
		changK=l+1
	}
	if changK<h.Size()&&(*h)[changK].count<(*h)[k].count{
		h.swap(changK,k)
		h.shiftDown(changK)
	}
}


func main() {
	topKFrequent([]int{1,1,1,2,2,3},2)
}
