package main

/**
198. 打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//尝试从nums[index...n]
func tryRob(nums []int,index int,memo []int) int  {
	if index>=len(nums){
		return 0
	}
	if memo[index]!=-1{
		return memo[index]
	}
	res:=0
	for i:=index;i<len(nums);i++{
		res=max(res,nums[i]+tryRob(nums,i+2,memo))
	}
	memo[index]=res
	return res
}




func rob(nums []int) int {
	n:=len(nums)
	//记忆化搜索
	memo:=make([]int,n)
	for i:=0;i<n;i++{
		memo[i]=-1
	}
	return tryRob(nums,0,memo)
}
func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func rob3(nums []int) int  {
	n:=len(nums)
	if n==0{
		return 0
	}
	dp:=make([]int,n+1)
	dp[0]=0
	dp[1]=nums[0]
	for i:=2;i<=n;i++{
		dp[i]=max(dp[i-1],nums[i-1]+dp[i-2])
	}
	return dp[n]
}

func rob2(nums []int) int  {
	//动态规划
	n:=len(nums)
	if n==0{
		return 0
	}
	memo:=make([]int,n)
	memo[n-1]=nums[n-1]
	for i:=n-2;i>=0;i--{
		for j:=i;j<n;j++{
			if j+2<n{
				memo[i]=max(memo[i],nums[j]+memo[j+2])
			}else{
				memo[i]=max(memo[i],nums[j])
			}
		}
	}
	return memo[0]
}

func main() {
	
}
