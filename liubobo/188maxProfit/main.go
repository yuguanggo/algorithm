package main

/**
188. 买卖股票的最佳时机 IV
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maxProfit(k int, prices []int) int {
	n:=len(prices)
	if n==0{
		return 0
	}
	dp:=make([][][2]int,n)
	if k>=n/2{
		res:=0
		for i:=1;i<n;i++{
			if prices[i]>prices[i-1]{
				res+=prices[i]-prices[i-1]
			}
		}
		return res
	}
	for i:=0;i<n;i++{
		dp[i]=make([][2]int,k+1)
	}
	for i:=0;i<n;i++{
		for j:=k;j>=1;j--{
			if i-1==-1{
				dp[i][j][0]=0
				dp[i][j][1]=-prices[i]
				continue
			}
			dp[i][j][0]=max(dp[i-1][j][0],dp[i-1][j][1]+prices[i])
			dp[i][j][1]=max(dp[i-1][j][1],dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}


func main() {
	
}
