package main

import "sort"

/**
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func permuteUnique(nums []int) [][]int {
	var res=make([][]int,0)
	var record=make(map[int]bool,0)
	sort.Ints(nums)
	helper(nums,0,[]int{},&res,record)
	return res
}

func helper(nums []int,index int,path []int,res *[][]int,record map[int]bool)  {
	if index==len(nums){
		var p=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=0;i<len(nums);i++{

		if !record[i]{
			if i>0&&nums[i-1]==nums[i]&&!record[i-1]{
				continue
			}
			record[i]=true
			path=append(path,nums[i])
			helper(nums,index+1,path,res,record)
			path=path[0:len(path)-1]
			record[i]=false
		}
	}
	return
}

func main() {
	
}
