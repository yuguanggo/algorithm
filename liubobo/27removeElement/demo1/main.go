package main

func removeElement(nums []int, val int) int {
	n:=len(nums)
	k:=0 //[o..k]为非0
	for i:=0;i<n;i++{
		if nums[i]!=val{
			if i!=k{
				nums[k],nums[i]=nums[i],nums[k]
				k++
			}else{
				k++
			}
		}
	}
	return k
}
func main() {
	
}
