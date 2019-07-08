package main

import "fmt"

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
	n:=len(nums)
	m:=make(map[int][][2]int)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{

			m[nums[i]+nums[j]]=[2]int{i,j}
		}
	}
	res:=[][]int{}
	visited:=make(map[int]bool)
	for i:=0;i<n;i++{
		complment:=0-nums[i]
		if v,ok:=m[complment];ok{
			if v[0]!=i&&v[1]!=i{
				indexsum:=v[0]+v[1]+i
				if _,ok:=visited[indexsum];!ok{
					res=append(res,[]int{nums[i],nums[v[0]],nums[v[1]]})
					visited[indexsum] = true
				}

			}
		}
	}
	return res
}
func main() {
	arr:=[]int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum(arr))
}
