package main

import "fmt"

/**
接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/

func trap(height []int) int {
	var (
		n        int
		max      int
		maxIndex int
		leftSum  int
		leftMax  int
		rightSum int
		rightMax int
	)
	n = len(height)
	//先找出最大值
	for i := 0; i < n; i++ {
		if height[i] > max {
			max = height[i]
			maxIndex = i
		}
	}
	//从左边遍历到最大值并计算雨水
	for i := 0; i < maxIndex; i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		} else {
			leftSum = leftSum + leftMax - height[i]
		}
	}
	//从右边遍历到最大值并计算雨水
	for j := n - 1; j > maxIndex; j-- {
		if height[j] > rightMax {
			rightMax = height[j]
		} else {
			rightSum = rightSum + rightMax - height[j]
		}
	}
	return leftSum + rightSum
}
func main() {
  var height = []int{0,1,0,2,1,0,1,3,2,1,2,1}
  fmt.Println(trap(height))
}
