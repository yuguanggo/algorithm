package main

func integerBreak(n int) int {
	dp:=make([]int,n+1)
	dp[1]=1
	for i:=2;i<=n;i++{
		for j:=1;j<=i-1;j++{
			dp[i]=max(dp[i],max(dp[i-j]*(i-j),j*(i-j)))
		}
	}
	return dp[n]
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
