package main

import "sort"

/**
合并区间
给出一个区间的集合，请合并所有重叠的区间。
示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 */
func merge(intervals [][]int) [][]int {
	var (
		n   int
		i   int
		res [][]int
	)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	n = len(intervals)
	i = 0
	for i < n {
		left := intervals[i][0]
		right := intervals[i][1]
		j := i + 1
		for j < n && intervals[j][0] <= right {
			if intervals[j][1]>right{
				right = intervals[j][1]
			}
			j++
		}
		res = append(res, []int{left, right})
		i = j
	}
	return res
}

func main() {

}
