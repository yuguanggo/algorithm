package main

//dp[i][0]=max(dp[i-1][0],dp[i-1][1]+price[i])
//dp[i][1]=max(dp[i-1][1],dp[i-2][0]-price[i])
func maxProfit(prices []int) int {
	n:=len(prices)
	if n==0{
		return 0
	}
	dp:=make([][2]int,n)
	for i:=0;i<n;i++{
		if i==0{
			dp[i][0]=0
			dp[i][1]=-prices[i]
		}else{
			dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
			if i>=2{
				dp[i][1]=max(dp[i-1][1],dp[i-2][0]-prices[i])
			}else{
				dp[i][1]=max(dp[i-1][1],-prices[i])
			}
		}
	}
	return dp[n-1][0]
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}
func main() {
	
}
