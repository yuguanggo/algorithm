package main

import (
	"fmt"
	"math"
	"sort"
)

/**
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return ans
			}

		}
	}
	return ans
}

func main() {
	fmt.Println(math.Abs(float64(-89)))
}
