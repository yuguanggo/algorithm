package main

import (
	"fmt"
	"sort"
)

/**
15. 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n:=len(nums)
	res:=[][]int{}
	for i:=0;i<n;i++{
		if nums[i]>0{
			break
		}
		if i>0&&nums[i] == nums[i-1] {
			continue
		}
		target:=0-nums[i]
		l:=i+1
		r:=n-1
		for l<r{
			if nums[l]+nums[r]==target{
				res = append(res,[]int{nums[i],nums[l],nums[r]})
				for l<r&&nums[l]==nums[l+1]{
					l++
				}
				for l<r&&nums[r]==nums[r-1]{
					r--
				}
				l++
				r--
			}else if nums[l]+nums[r]<target{
				l++
			}else{
				r--
			}
		}
	}
	return res
}


func main() {
	arr:=[]int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum(arr))
}
