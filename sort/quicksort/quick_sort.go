package quicksort

//对arr[l..r]进行partition操作
//使得arr[l..p-1]<arr[p] arr[p+1..r]>arr[p]
func partition(arr []int, l int, r int) int {
	//此处对于近乎有序数列会退化成o(n2) 过多重复键值也会使快速排序退化成o(n2)
	//var v = arr[l]

	//进行优化，随机选择一个标定点
	//var p = rand.Intn(r-l)+l
	//
	//arr[l],arr[p] = arr[p],arr[l]
	//var v = arr[l]
	arr[l],arr[(l+r)/2] = arr[(l+r)/2],arr[l]
	var v = arr[l]
	var j = l //一直指向小于v的最末尾
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] =arr[i],arr[j]
		}
	}
	arr[j], arr[l] = arr[l],arr[j]
	return j
}
// 递归使用快速排序,对arr[l...r]的范围进行排序
func sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	var p = partition(arr, l, r)
	//此处去掉p
	sort(arr, l, p-1)
	sort(arr, p+1, r)
}

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}
