package maxheap

import (
	"errors"
)

type MaxHeap []int

func (h MaxHeap) shiftUp(k int) {
	for k > 0 && h[(k-1)/2] < h[k] {
		h.Swap(k, (k-1)/2)
		k = (k - 1) / 2
	}
}

func (h MaxHeap) shiftDown(k int) {
	l := h.Size()
	for (2*k + 1) < l {
		j := 2*k + 1 //左子树
		if (j+1) < l && h[j] < h[j+1] {
			j++
		}
		if h[k] >= h[j] {
			break
		}
		h.Swap(k, j)
		k = j
	}
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Size() int {
	return len(h)
}

func (h MaxHeap) IsEmpty() bool {
	return len(h) == 0
}

func (h *MaxHeap) Insert(v int) {
	*h = append(*h, v)
	h.shiftUp(h.Size() - 1)
}

//将一个数组堆化
func (h *MaxHeap) MaxHeap(arr []int) {
	n := len(arr)
	*h = arr
	for i := (n - 2) / 2; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *MaxHeap) ExtractMax() (int, error) {
	old := *h
	if len(old) == 0 {
		return 0, errors.New("slice is empty")
	}
	v := old[0]
	old[0], old[len(old)-1] = old[len(old)-1], old[0]
	*h = old[:len(old)-1]
	h.shiftDown(0)
	return v, nil
}

func HeapSort(arr []int) {
	var h MaxHeap
	h.MaxHeap(arr)
	n := h.Size()
	for i := n - 1; i >= 0; i-- {
		v, err := h.ExtractMax()
		if err == nil {
			arr[i] = v
		}
	}
}

// heapSort2, 借助我们的heapify过程创建堆
// 此时, 创建堆的过程时间复杂度为O(n), 将所有元素依次从堆中取出来, 实践复杂度为O(nlogn)
// 堆排序的总体时间复杂度依然是O(nlogn), 但是比上述heapSort1性能更优, 因为创建堆的性能更优
func HeapSort2(arr []int) {
	//建立最大堆
	n := len(arr)
	if n == 0 {
		return
	}
	for i := (n - 2) / 2; i >= 0; i-- {
		//执行shiftdown操作
		k := i;
		for (2*k + 1) < n {
			j := 2*k + 1 //左子树
			//找出左右子树哪个值大
			if j+1 < n && arr[j] < arr[j+1] {
				j++
			}
			//开始和最大值比较
			if arr[k] > arr[j] {
				//父节点大于子节点不用换
				break
			}
			arr[k], arr[j] = arr[j], arr[k]
			k = j
		}
	}
	//从建好的堆中依次取出最大值
	var temp []int
	m := n
	for i := 0; i < m; i++ {
		//取出首个最大值
		v := arr[0]
		temp = append(temp,v)
		//和最后一个交换下位置
		arr[0], arr[n-1] = arr[n-1], arr[0]
		arr = arr[:n-1] //删除最后一个元素
		n = len(arr)
		//将第一元素执行shiftdown操作
		k := 0
		for 2*k+1 < n {
			j := 2*k + 1
			if j+1 < n && arr[j] < arr[j+1] {
				j++
			}
			if arr[k] > arr[j] {
				break
			}
			arr[k], arr[j] = arr[j], arr[k]
			k = j
		}
	}
	arr = temp
}

func shiftDown(i int,n int,arr []int)  {

}

func HeapSort3(arr []int)  {
	n:= len(arr)
	//开始建堆
	for i:=(n-2)/2;i>=0;i--{
		shiftDown(i,n,arr)
	}
}
