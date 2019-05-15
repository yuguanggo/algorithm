package main

import "fmt"

/**
搜索旋转排序数组
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别
 */
func search(nums []int, target int) int {
	var (
		left  int
		right int
		mid   int
	)
	left = 0
	right = len(nums) - 1

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid-1
			}
		}else{
			if nums[mid]>target&&nums[left]<=target{
				right = mid-1
			}else{
				left=mid+1
			}
		}
	}
	return -1
}

func main() {
 var num = []int{4,5,6,7,0,1,2}
 fmt.Println(search(num,0))
}
