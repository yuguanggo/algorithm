package main


/**
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func combinationSum3(k int, n int) [][]int {
	var res=make([][]int,0)
	helper(k,n,[]int{},&res,1)
	return res
}

func helper(k int,n int,path []int,res *[][]int,index int)  {
	if k==len(path)&&n==0{
		var p=make([]int,k)
		copy(p,path)
		*res=append(*res,p)
	}
	for i:=index;i<=9;i++{
		if i>n{
			break
		}
		path=append(path,i)
		helper(k,n-i,path,res,i+1)
		path=path[0:len(path)-1]
	}
}

func main() {
	
}
