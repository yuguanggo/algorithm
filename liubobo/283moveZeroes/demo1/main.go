package main

func moveZeroes(nums []int)  {
	n:=len(nums)
	k:=0 //[0..k]非0元素
	for i:=0;i<n;i++{
		if nums[i]!=0{
			if i!=k{
				nums[k],nums[i]=nums[i],nums[k]
				k++
			}else{
				k++
			}
		}
	}
}
func main() {
	
}
