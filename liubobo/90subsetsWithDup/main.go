package main

import "sort"

/**
90. 子集 II
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func subsetsWithDup(nums []int) [][]int {
	var res =make([][]int,0)
	sort.Ints(nums)
	helper(nums,0,[]int{},&res)
	return res
}

func helper(nums []int,index int,path []int,res *[][]int)  {
	var p=make([]int,len(path))
	copy(p,path)
	*res=append(*res,p)
	for i:=index;i<len(nums);i++{
		if i>index&&nums[i-1]==nums[i]{
			continue
		}
		path=append(path,nums[i])
		helper(nums,i+1,path,res)
		path=path[0:len(path)-1]
	}
}

func main() {
	
}
