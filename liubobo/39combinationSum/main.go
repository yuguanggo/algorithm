package main

import "fmt"

/**
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func combinationSum(candidates []int, target int) [][]int {
	var res=make([][]int,0)
	helper(candidates,target,0,0,[]int{},&res)
	return res
}

func helper(candidates []int,target int,sum int,index int,path []int,res *[][]int)  {
	if target==sum{
		var p=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=index;i<len(candidates);i++ {
		if sum>target{
			return
		}
		path=append(path,candidates[i])
		sum=sum+candidates[i]
		helper(candidates,target,sum,i,path,res)
		sum=sum-path[len(path)-1]
		path=path[0:len(path)-1]
	}
	return
}

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7},7))
}
