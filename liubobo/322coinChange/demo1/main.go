package main

import "math"

func coinChange(coins []int, amount int) int {
	dp:=make([]int,amount+1)
	dp[0]=0
	for i:=1;i<=amount;i++{
		dp[i]=math.MaxInt32
	}
	for i:=1;i<=amount;i++{
		for j:=0;j<len(coins);j++{
			if i>=coins[j]{
				dp[i]=min(dp[i],dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount]==math.MaxInt32{
		return -1
	}
	return dp[amount]
}

func min(i,j int) int  {
	if i>j{
		return j
	}
	return i
}



func main() {
	
}
