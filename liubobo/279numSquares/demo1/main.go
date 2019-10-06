package main

import (
	"math"
)

func numSquares(n int) int {
	//转换成图
	q:=make([][2]int,0)
	q=append(q,[2]int{n,0})//数组第2个表示当前是第几步
	visited:=make([]bool,n+1)
	visited[n]=true
	for len(q)>0{
		num:=q[0]
		q=q[1:]
		m:=num[0]
		step:=num[1]
		for i:=1;;i++{
			a:=m-i*i
			if a<0{
				break
			}
			if a==0{
				return step+1
			}
			if !visited[a]{
				q=append(q,[2]int{a,step+1})
				visited[a]=true
			}
		}
	}
	return 0
}
func numSquares3(n int) int {
	dp:=make([]int,n+1)
	for i:=1;i<=n;i++{
		dp[i]=i
		for j:=1;j*j<i;j++{
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
func numSquares2(n int) int {
	memo:=make([]int,n+1)
	return helper(n,memo)
}

func helper(n int,memo []int) int  {
	if memo[n]!=0{
		return memo[n]
	}
	for i:=1;i*i<=n;i++{
		if i*i==n{
			memo[n]=1
			return memo[n]
		}
	}

	res:=math.MaxInt32
	for i:=1;i*i<n;i++{
		res=min(res,helper(n-i*i,memo)+1)
	}
	memo[n]=res
	return memo[n]
}






func main() {
	
}
