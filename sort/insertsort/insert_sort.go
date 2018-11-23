package insertsort

func InsertSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func InsertSort1(arr []int, n int) {
	for i := 1; i < n; i++ {

		//要插入的元素
		item := arr[i]
		var j = i
		for ; j > 0; j-- {
			if item < arr[j-1] {
				arr[j] = arr[j-1]
			} else {
				break;
			}
		}
		arr[j] = item;
	}
}
