package main

import "fmt"

/**
最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 */
func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

/**
分治法
 */
func maxSubArray2(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) int {
	var (
		mid, lmax, rmax, lsum, rsum, lmaxsub, rmaxsub, max int
	)
	if left == right {
		return nums[left]
	}
	mid = (left + right) / 2
	lmax = helper(nums, left, mid)
	rmax = helper(nums, mid+1, right)
	lmaxsub=nums[mid]
	for i := mid; i >= left; i-- {
		lsum += nums[i]
		if lmaxsub < lsum {
			lmaxsub = lsum
		}
	}
	rmaxsub=nums[mid+1]
	for i := mid + 1; i <= right; i++ {
		rsum += nums[i]
		if rmaxsub < rsum {
			rmaxsub = rsum
		}
	}
	max = lmax
	if rmax > max {
		max = rmax
	}
	if (lmaxsub + rmaxsub) > max {
		max = lmaxsub + rmaxsub
	}
	return max
}

func main() {
	s := []int{-2, -1}
	fmt.Println(maxSubArray2(s))
}
