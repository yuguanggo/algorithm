package main

func moveZeroes(nums []int) {
	n:=len(nums)
	j:=0 //[0..j]为非0
	for i:=0;i<n;i++{
		if nums[i]!=0{
			if i!=j{
				nums[i],nums[j]=nums[j],nums[i]
				j++
			}else{
				j++
			}
		}
	}

}
func main() {
	
}
