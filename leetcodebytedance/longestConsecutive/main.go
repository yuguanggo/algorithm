package main

import "fmt"

/**
 最长连续序列
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
用哈希表存储每个端点值对应连续区间的长度
若数已在哈希表中：跳过不做处理
若是新数加入：
取出其左右相邻数已有的连续区间长度 left 和 right
计算当前数的区间长度为：cur_length = left + right + 1
根据 cur_length 更新最大长度 max_length 的值
更新区间两端点的长度值
 */
func longestConsecutive(nums []int) int {
	var (
		max    int
		left   int
		right  int
		length int
		m      map[int]int
	)
	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			left = m[nums[i]-1]
			right = m[nums[i]+1]
			length = left + right + 1
			if max < length {
				max = length
			}
			m[nums[i]]=length
			m[nums[i]-left] =length
			m[nums[i]+right] = length
		}
	}
	return max
}

func main() {
	var nums = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))

}
