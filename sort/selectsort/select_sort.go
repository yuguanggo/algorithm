package selectsort

func SelectSort(arr []int,n int){
	for i := 0; i < n; i++ {
		//最小值所在的索引
		minIndex:=i
		for j := i+1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}
