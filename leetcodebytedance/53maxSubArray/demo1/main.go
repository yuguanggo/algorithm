package main

func maxSubArray(nums []int) int {
	ans:=nums[0]
	sum:=0
	for i:=0;i<len(nums);i++{
		if sum>0{
			sum+=nums[i]
		}else{
			sum=nums[i]
		}
		if ans<sum{
			ans=sum
		}
	}
	return ans
}


func main() {
	
}
