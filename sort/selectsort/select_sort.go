package selectsort

func SelectSort(arr []int) {
	n :=len(arr)
	for i := 0; i < n; i++ {
		//最小值所在的索引
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func SelectSort1(arr []int) {
	n :=len(arr)
	left := 0
	right := n - 1
	for left < right {
		minIndex := left
		maxIndex := right

		//每轮查找是要保证 arr[minIndex]<=arr[maxIndex]
		if arr[minIndex] > arr[maxIndex] {
			arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
		}

		for i := left + 1; i < right; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			} else if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		arr[left],arr[minIndex] = arr[minIndex],arr[left]
		arr[right],arr[maxIndex] = arr[maxIndex],arr[right]
		left++
		right--
	}
}
