package main

import "fmt"

/**
77. 组合

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func combine(n int, k int) [][]int {
	var res =make([][]int,0)
	helper(n,k,1,[]int{},&res)
	return res
}

func helper(n int,k int,index int,path []int,res *[][]int)  {
	if len(path)==k{
		var p=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=index;i<=n-(k-len(path))+1;i++{
		fmt.Println("i:",i)
		path=append(path,i)
		helper(n,k,i+1,path,res)
		path=path[0:len(path)-1]
	}
	return
}

func main() {
	fmt.Println(combine(4,2))
}
