package main

//dp[i]=dp[i-nums[0]]+dp[1-nums[1]]+...+dp[i-nums[len(nums)-1]]
func combinationSum4(nums []int, target int) int {
	dp:=make([]int,target+1)
	dp[0]=1
	for i:=1;i<=target;i++{
		for j:=0;j<len(nums);j++{
			if i>=nums[j]{
				dp[i]=dp[i]+dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}



func main() {
	
}
