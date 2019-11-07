package main


//dp[i]=max(nums[i],dp[i-1]+nums[i])
func maxSubArray(nums []int) int {
	n:=len(nums)
	pre:=nums[0]
	maxSum:=pre
	for i:=1;i<n;i++{
		sum:=max(nums[i],pre+nums[i])
		maxSum=max(sum,maxSum)
		pre=sum
	}
	return maxSum
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
