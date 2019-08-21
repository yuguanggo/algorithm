package main

import "sort"

/**
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func combinationSum2(candidates []int, target int) [][]int {
	var res=make([][]int,0)
	sort.Ints(candidates)
	helper(candidates,target,0,[]int{},&res)
	return res
}

func helper(candidates []int, target int,index int,path []int,res *[][]int)  {
	if target==0{
		var p=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=index;i<len(candidates);i++{
		if candidates[i]>target{
			break
		}
		if i>index&&candidates[i-1]==candidates[i]{
			continue
		}
		path=append(path,candidates[i])
		helper(candidates,target-candidates[i],i+1,path,res,)
		path=path[0:len(path)-1]
	}
}


func main() {
	
}
