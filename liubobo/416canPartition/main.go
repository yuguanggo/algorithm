package main

/**
416. 分割等和子集
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].
 

示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func canPartition(nums []int) bool {
	sum:=0
	n:=len(nums)
	for i:=0;i<n;i++{
		sum+=nums[i]
	}
	if sum%2!=0 {
		return false
	}
	c:=sum/2
	dp:=make([]bool,n+1)
	//初始化第一行
	for i:=0;i<=c;i++{
		dp[i]= nums[0]==i
	}
	for i:=1;i<n;i++{
		for j:=c;j>=nums[i];j--{
			dp[j]=dp[j]||dp[j-nums[i]]
		}
	}
	return dp[c]
}


func main() {
	
}
