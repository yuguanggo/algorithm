package main

import "math"

/**
121. 买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析
一共有三种状态
i 表示第几天
k 表示交易几次
0,1表示没有持有或者持有股票
状态转移
0 买入变成1  1卖出变成0
0 不操作 0 1不操作1
dp[i][k][0]表示第i天至多交易了k次没有持有股票
dp[i][k][1]表示第i天至多交易了k次持有股票
最大利润为dp[i][k][0] 没卖出的话不算利润
状态转移方程
dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+price[i])
第i天没有持有股票，有两种可能
昨天就没有持有股票，或者是昨天持有股票，但今天卖了
dp[i][k][1]=max(dp[i-1][k][1],dp[k-1][0]-price[i])
第i天持有股票，有两种可能
昨天就持有股票，或者昨天没有持有，但今天买入了
解决第一题
k限制成1 所以k没用了
 */
func maxProfit2(prices []int) int {
	n:=len(prices)
	if n==0{
		return 0
	}
	dp:=make([][2]int,n)
	for i:=0;i<n;i++{
		if i-1==-1{
			dp[i][0]=0
			dp[i][1]=-prices[i]
			continue
		}
		dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1]=max(dp[i-1][1],-prices[i])
	}
	return dp[n-1][0]
}

func maxProfit(prices []int) int {
	//简化 o(1)
	n:=len(prices)
	if n==0{
		return 0
	}
	dp_i_0:=0
	dp_i_1:=math.MinInt32
	for i:=0;i<n;i++{
		dp_i_0=max(dp_i_0,dp_i_1+prices[i])
		dp_i_1=max(dp_i_1,-prices[i])
	}
	return dp_i_0
}


func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}
func main() {
	
}
