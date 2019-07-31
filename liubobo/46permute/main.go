package main

import (
	"fmt"
)

/**
46. 全排列
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func permute(nums []int) [][]int {
	var res=make([][]int,0)
	var record=make(map[int]bool,0)
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
	m:=make(map[int]bool,0)
	fmt.Println(!m[1])
}
