package main


func canPartition(nums []int) bool {
	n:=len(nums)
	sum:=0
	for i:=0;i<n;i++{
		sum+=nums[i]
	}
	if sum%2!=0{
		return false
	}
	sum=sum/2
	dp:=make([]bool,sum+1)
	for i:=0;i<=sum;i++{
		if nums[0]==i{
			dp[i]=true
			break
		}
	}
	for i:=1;i<n;i++{
		for j:=sum;j>=nums[i];j--{
			dp[j]=dp[j]||dp[j-nums[i]]
		}
	}
	return dp[sum]
}
func main() {
	
}
