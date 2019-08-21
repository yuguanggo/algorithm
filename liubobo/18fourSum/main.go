package main

import "sort"

/**
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n:=len(nums)
	res:=[][]int{}
	for i:=0;i<n;i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		for j:=i+1;j<n;j++{
			if j>i+1&&nums[j]==nums[j-1]{
				continue
			}
			l:=j+1
			r:=n-1
			target2:=target-nums[i]-nums[j]
			for l<r{
				if nums[l]+nums[r]==target2{
					res = append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
					for l<r&&nums[l]==nums[l+1]{
						l++
					}
					for l<r&&nums[r]==nums[r-1]{
						r--
					}
					l++
					r--
				}else if nums[l]+nums[r]<target2{
					l++
				}else{
					r--
				}
			}
		}
	}
	return res
}

func main() {
	
}
