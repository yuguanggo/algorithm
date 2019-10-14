package main

func rob(nums []int) int {
	n:=len(nums)
	if n==1{
		return nums[0]
	}
	dpa:=make([]int,n+1)
	dpb:=make([]int,n+1)
	for i:=2;i<=n;i++{
		//偷第一个
		dpa[i]=max(dpa[i-1],dpa[i-2]+nums[i-2])
		//偷最后一个
		dpb[i]=max(dpb[i-1],dpb[i-2]+nums[i-1])
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
