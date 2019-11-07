package main

import "fmt"

/**
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
 /**
 查询左右边界
  */
func binarySearch(nums []int, target int,lorR bool) int  {
	l:=0
	r:=len(nums)-1
	for l<=r{
		mid:=l+(r-l)/2
		if nums[mid]==target{
			if lorR{
				//左边界
				r=mid-1
			}else{
				l=mid+1
				//右边界
			}
		}else if nums[mid]>target{
			r=mid-1
		}else {
			l=mid+1
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	res:=[]int{-1,-1}
	if len(nums)==0{
		return res
	}
	lindex:=binarySearch(nums,target,true)
	if lindex==len(nums)|| nums[lindex]!=target{
		return res
	}
	rindex:=binarySearch(nums,target,false)
	res=[]int{lindex,rindex-1}
	return res
}


func main() {
	nums:=[]int{2,2}
	//fmt.Println(searchRange(nums,8))
	fmt.Println(binarySearch(nums,3,true))
	fmt.Println(binarySearch(nums,3,false))
}
