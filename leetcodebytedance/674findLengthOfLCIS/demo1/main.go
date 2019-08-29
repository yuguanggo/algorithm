package main


func findLengthOfLCIS(nums []int) int {
	if len(nums)==0{
		return 0
	}
	i:=0
	res:=1
	for j:=1;j<len(nums);j++{
		if nums[j]>nums[j-1]{
			res=max(res,j-i+1)
		}else{
			i=j
		}
	}
	return res
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
