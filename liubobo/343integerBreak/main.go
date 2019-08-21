package main


/**
343. 整数拆分
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func integerBreak2(n int) int  {
	//自顶向下的记忆化搜索
	memo:=make([]int,n+1)
	memo[0]=0
	memo[1]=1
	for i:=2;i<=n;i++{
		memo[i]=-1
	}
	return breakInteger(n,memo)
}

func breakInteger(n int,memo []int) int {
	if memo[n]!=-1{
		return memo[n]
	}
	res:=0
	for i:=1;i<=n-1;i++{
		res=max3(res,i*(n-i),i*breakInteger(n-i,memo))
	}
	memo[n]=res
	return res
}

func integerBreak(n int) int {
	//自底向上的动态规划
	memo:=make([]int,n+1)
	memo[1]=1
	for i:=2;i<=n;i++{
		for j:=1;j<=i-1;j++{
			memo[i]=max3(memo[i],j*(i-j),j*memo[i-j])
		}
	}
	return memo[n]
}

func max3(i,j,k int) int  {
	return max(i,max(j,k))
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}




func main() {
	
}
