package main


/**
213. 打家劫舍 II
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [2,3,2]
输出: 3
解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2:

输入: [1,2,3,1]
输出: 4
解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func rob(nums []int) int {
	n:=len(nums)
	if n==1{
		return nums[0]
	}
	dpa:=make([]int,n+1)
	dpb:=make([]int,n+1)
	for i:=2;i<=n;i++{
		//去掉第一个
		dpa[i]=max(dpa[i-1],dpa[i-2]+nums[i-1])
		//去掉最后一个
		dpb[i]=max(dpb[i-1],dpb[i-2]+nums[i-2])
	}
	return max(dpa[n],dpb[n])
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
