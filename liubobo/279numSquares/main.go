package main

import (
	"fmt"
)

/**
279. 完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func numSquares(n int) int {
	//转换为图
	queue:=make([][2]int,0)
	queue=append(queue,[2]int{n,0})
	visited:=make([]bool,n+1)
	visited[n]=true
	for len(queue)>0{
		fornt:=queue[0]
		num:=fornt[0]
		step:=fornt[1]
		queue=queue[1:]
		for i:=1;;i++{
			a:=num-i*i
			if a<0{
				break
			}
			if a==0{
				return step+1
			}
			if !visited[a]{
				queue=append(queue,[2]int{a,step+1})
				visited[a]=true
			}
		}
	}
	return 0
}

func numSquares2(n int) int {
	//动态规划
	dp:=make([]int,n+1)
	for i:=1;i<=n;i++{
		dp[i]=i
		for j:=1;i-j*j>=0;j++{
			dp[i]=min(dp[i],dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(i,j int) int  {
	if i>j{
		return j
	}
	return i
}

func main() {
	fmt.Println(numSquares2(12))
}
