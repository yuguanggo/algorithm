package shell_sort

func ShellSort(arr []int) {
	n := len(arr)
	//设置增量
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			//对arr[i],arr[i-h],arr[i-2h]进行插入排序
			item:=arr[i]
			j := i
			for ; j >= h; j -= h {
				if arr[j] < arr[j-h] {
					arr[j]=arr[j-h]
				}
			}
			arr[j] = item
		}
		h /= 3
	}
}
