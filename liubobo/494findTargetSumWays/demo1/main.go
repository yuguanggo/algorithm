package main


/**
sum(p)-sum(n)=S
sum(p)+sum(n)+sum(p)-sum(n)=S+sum(p)+sum(n)
2sum(p)=S+sum(nums)
sum(p)=(S+sum(nums))/2
 */
func findTargetSumWays(nums []int, S int) int {
	sum:=0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	if sum<S{
		return 0
	}
	if (S+sum)%2!=0{
		return 0
	}
	p:=(sum+S)/2
	dp:=make([]int,p+1)
	dp[0]=1
	for i:=0;i<len(nums);i++{
		for j:=p;j>=nums[i];j--{
			dp[j]=dp[j]+dp[j-nums[i]]
		}
	}
	return dp[p]
}


func main() {
	
}
