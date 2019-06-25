package quicksort

//对arr[l..r]进行partition操作
//使得arr[l..p-1]<arr[p] arr[p+1..r]>arr[p]
func partition(arr []int, l int, r int) int {
	//此处对于近乎有序数列会退化成o(n2)
	//var v = arr[l]

	//进行优化，随机选择一个标定点
	//var p = rand.Intn(r-l)+l
	//
	//arr[l],arr[p] = arr[p],arr[l]
	//var v = arr[l]
	arr[l], arr[(l+r)/2] = arr[(l+r)/2], arr[l]
	var v = arr[l]
	var j = l //一直指向小于v的最末尾
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}

//对arr[l..r]进行partition操作
//使得arr[l..i-1]<arr[p] arr[j+1..r]>arr[p]
//func partition2(arr []int, l int, r int) int {
//	i := l + 1 //左边排序的起始点
//	j := r     //右边排序的起始点
//	v := arr[l]
//	for i <= j {
//		if arr[i] > v {
//			arr[i],arr[j] = arr[j],arr[i]
//			j--
//		}else{
//			i++
//		}
//		if arr[j]<v{
//			arr[i],arr[j] = arr[j],arr[i]
//			i++
//		}else{
//			j--
//		}
//	}
//	return j
//}

// 双路快速排序的partition 解决过多重复键值也会使快速排序退化成o(n2)
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition2(arr []int, l int, r int) int {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	arr[l], arr[(l+r)/2] = arr[(l+r)/2], arr[l]
	v := arr[l]
	// arr[l+1...i) <= v; arr(j...r] >= v
	i := l + 1
	j := r
	for {
		// 注意这里的边界, arr[i].compareTo(v) < 0, 不能是arr[i].compareTo(v) <= 0
		// 思考一下为什么?
		for i <= r && arr[i] < v {
			i++
		}
		// 注意这里的边界, arr[j].compareTo(v) > 0, 不能是arr[j].compareTo(v) >= 0
		// 思考一下为什么?
		for j >= l+1 && arr[j] > v {
			j--
		}
		// 对于上面的两个边界的设定, 有的同学在课程的问答区有很好的回答:)
		// 大家可以参考: http://coding.imooc.com/learn/questiondetail/4920.html
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
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

//双路快排
func sort2(arr []int, l int, r int) {
	if l >= r {
		return
	}
	var p = partition2(arr, l, r)
	//此处去掉p
	sort2(arr, l, p-1)
	sort2(arr, p+1, r)
}

//三路快排
func sort3(arr []int, l int, r int) {
	if l >= r {
		return
	}

	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	arr[l], arr[(l+r)/2] = arr[(l+r)/2], arr[l]
	v := arr[l]

	// arr[l..lt-1]<v arr[lt...gt-1] ==v arr[jt+1]>v
	lt := l     //arr[l..lt]<v
	gt := r + 1 //arr[gt...r]>v
	i := l + 1  // arr[lt+1..i]==v
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	sort3(arr, l, lt-1)
	sort3(arr, gt, r)

}

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func QuickSort2(arr []int) {
	sort2(arr, 0, len(arr)-1)
}

func QuickSort3(arr []int) {
	sort3(arr, 0, len(arr)-1)
}
