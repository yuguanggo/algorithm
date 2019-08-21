package main

import "fmt"

/**
347. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func topKFrequent(nums []int, k int) []int {
	records:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		records[nums[i]]++
	}
	var heap MinHeap
	for num,count:=range records{
		if heap.Size()>=k{
			top:=heap.Min()
			if top.count<count{
				heap.ExtraMin()
			}
		}
		if heap.Size()<k{
			heap.Insert(node{num,count})
		}
	}
	res:=make([]int,k)
	for i:=k-1;i>=0;i--{
		res[i]=heap.ExtraMin().v
	}
	return res
}

type node struct {
	v int
	count int
}

type MinHeap []node

func (h *MinHeap)Size() int {
	return len(*h)
}

func (h *MinHeap)Insert(n node)  {
	*h = append(*h,n)
	h.shiftUp(h.Size()-1)
}
func (h *MinHeap)Swap(i,j int)  {
	(*h)[i],(*h)[j]=(*h)[j],(*h)[i]
}
func (h *MinHeap)shiftUp(k int)  {
	for k>0&&(*h)[(k-1)/2].count>(*h)[k].count{
		h.Swap(k,(k-1)/2)
		k=(k-1)/2
	}
}

func (h *MinHeap)shiftDown(k int)  {
	l:=h.Size()
	for (2*k+1)<l{
		j:=2*k+1 //左子树
		if j+1<l&&(*h)[j].count>(*h)[j+1].count{
			j++
		}
		if (*h)[k].count<(*h)[j].count{
			break
		}
		h.Swap(k,j)
		k=j
	}
}

func (h *MinHeap)Min() node  {
	if h==nil{
		return node{}
	}
	return (*h)[0]
}

func (h *MinHeap)ExtraMin() node  {
	if len(*h)==0{
		return node{}
	}
	n:=(*h)[0]
	h.Swap(0,h.Size()-1)
	(*h)=(*h)[:h.Size()-1]
	h.shiftDown(0)
	return n
}

func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3},2))

}
