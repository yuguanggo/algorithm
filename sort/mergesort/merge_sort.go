package mergesort

import (
	"algorithm/sort/insertsort"
	"math"
)

func merge(arr []int, left int, middle int, right int) {
	var size = right - left + 1
	var cparr = make([]int, size)
	for m := left; m <= right; m++ {
		cparr[m-left] = arr[m]
	}
	// 初始化，i指向左半部分的起始索引位置l；j指向右半部分起始索引位置mid+1
	i, j := left, middle+1
	for k := left; k <= right; k++ {
		if i > middle {
			//左边边已经排序完
			arr[k] = cparr[j-left]
			j++
		} else if j > right {
			arr[k] = cparr[i-left]
			i++
		} else if cparr[i-left] < cparr[j-left] {
			arr[k] = cparr[i-left]
			i++
		} else {
			arr[k] = cparr[j-left]
			j++
		}
	}
}

func sort(arr []int, left int, right int) {
	//if left>=right{
	//	return
	//}
	if right-left <= 15 {
		insertsort.InsertSort2(arr, left, right)
		return
	}
	middle := (left + right) / 2
	sort(arr, left, middle)
	sort(arr, middle+1, right)
	// 对于arr[mid] <= arr[mid+1]的情况,不进行merge
	// 对于近乎有序的数组非常有效,但是对于一般情况,有一定的性能损失
	if arr[middle] > arr[middle+1] {
		merge(arr, left, middle, right)
	}

}

func MergeSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func MergeSortBU(arr []int) {
	// Merge Sort BU 也是一个O(nlogn)复杂度的算法，虽然只使用两重for循环
	// 所以，Merge Sort BU也可以在1秒之内轻松处理100万数量级的数据
	// 注意：不要轻易根据循环层数来判断算法的复杂度，Merge Sort BU就是一个反例
	n := len(arr)
	//无优化版本
	//for sz := 1; sz < n; sz *= 2 {
	//	for i := 0; i < n-sz; i += sz + sz {
	//		merge(arr,i,i+sz-1,int(math.Min(float64(i+sz+sz-1),float64(n-1))))
	//	}
	//}
	//优化版本
	for i:=0;i<n;i+=16{
		//对于小于16位的数组使用插入排序
		insertsort.InsertSort2(arr,i,int(math.Min(float64(i+15),float64(n-1))))
	}


	for sz:=16;sz<n;sz*=2{
		for i:=0;i<n-sz;i+=sz+sz{
			if arr[i+sz-1]>arr[i+sz]{
				//对于arr[mid]<arr[mid+1]的情况不排序
				merge(arr,i,i+sz-1,int(math.Min(float64(i+2*sz-1),float64(n-1))))
			}
		}
	}

}
