package main

import "fmt"

/**
120. 三角形最小路径和
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func minimumTotal(triangle [][]int) int {
	n:=len(triangle)
	dp:=make([]int,n)
	for i:=n-1;i>=0;i--{
		for j:=0;i<len(triangle[n-1]);j++{
			if i==n-1{
				dp[j]=triangle[i][j]
			}else{
				dp[j]=min(dp[j],dp[j+1])+triangle[i][j]
			}
		}
	}
	return dp[0]
}

func min(i,j int) int  {
	if i>j{
		return j
	}
	return i
}
func main() {
	t:=[][]int{{-1},{2,3},{1,-1,-3}}
	fmt.Println(minimumTotal(t))
}
