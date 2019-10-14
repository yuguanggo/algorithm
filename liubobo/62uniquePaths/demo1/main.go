package main

func uniquePaths(m int, n int) int {
	dp:=make([]int,n+1)
	for i:=1;i<=m;i++{
		for j:=1;j<n;j++{
			if i==1{
				//初始化第一行
				dp[j]=1
			}else{
				dp[j]=dp[j-1]+dp[j]
			}
		}
	}
	return dp[n]
}

func main() {
	
}
