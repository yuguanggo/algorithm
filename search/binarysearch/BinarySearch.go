package binarysearch

//非递归二分查找
func Find(arr []int, v int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == v {
			return mid
		}

		if arr[mid] > v {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func RecursionFind(arr []int, l, r, v int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if arr[mid] == v {
		return mid
	}
	if arr[mid] > v {
		return RecursionFind(arr, l, mid-1, v)
	} else {
		return RecursionFind(arr, mid+1, r, v)
	}
}
