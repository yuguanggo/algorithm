package main

import (
	"fmt"
	"algorithm/liubobo/215findKthLargest/quicksort"
)

/**
数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
 //进行三路快排
 //nums[l..lt]>v nums[lt+1..i-1]=v nums[gt..r]<v
func quickSort(nums []int,l,r,k int) int  {
	mid:=l+(r-l)/2
	nums[l],nums[mid] = nums[mid],nums[l]
	v:=nums[l]
	i:=l+1
	lt:=l
	gt:=r+1
	for i<gt{
		if nums[i]>v{
			nums[i],nums[lt+1]=nums[lt+1],nums[i]
			lt++
			i++
		}else if nums[i]<v{
			nums[i],nums[gt-1]=nums[gt-1],nums[i]
			gt--
		}else{
			i++
		}
	}
	nums[l],nums[lt]=nums[lt],nums[l]
	if k-1<=lt-1{
		return quickSort(nums,l,lt-1,k)
	}else if k-1>=gt{
		return quickSort(nums,gt,r,k)
	}else{
		return nums[k-1]
	}
}
func findKthLargest(nums []int, k int) int {
	return quickSort(nums,0,len(nums)-1,k)
}

func main() {
	var arr =[]int{3,2,1,5,6,4}
	quicksort.ReSort3(arr,0,len(arr)-1)
	fmt.Println(arr)
}
