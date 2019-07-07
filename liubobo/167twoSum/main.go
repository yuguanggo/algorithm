package main

import (
	"fmt"
)

/**
两数之和 II - 输入有序数组
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 */
 //在numbers[l...r] 查找等于k的元素索引
func binarySearch(numbers []int,l,r,k int) int  {
	if l>r{
		return -1
	}
	mid:=l+(r-l)/2
	if numbers[mid]<k{
		return binarySearch(numbers,mid+1,r,k)
	}else if numbers[mid]>k{
		return binarySearch(numbers,l,mid-1,k)
	}else {
		return mid
	}
}
func twoSumBinarySearch(numbers []int, target int) []int {
	n:=len(numbers)
	for i:=0;i<n;i++{
		p:=binarySearch(numbers,i+1,n-1,target-numbers[i])
		if(p!=-1){
			return []int{i+1,p+1}
		}
	}
	return []int{-1,-1}
}

func twoSum(numbers []int, target int) []int {
	n:=len(numbers)
	i:=0
	j:=n-1
	for i<j{
		if numbers[i]+numbers[j]>target{
			j--
		}else if numbers[i]+numbers[j]<target{
			i++
		}else{
			return []int{i+1,j+1}
		}
	}
	return []int{0,0}

}


func main() {
	var arr =[]int{2, 7, 11, 15}
	twoSum(arr,9)
	fmt.Println(twoSum(arr,9))
}
