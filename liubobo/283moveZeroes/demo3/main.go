package main

func moveZeroes(nums []int) {

	index:=-1
	if nums[0]==0{
		index=0
	}
	for i:=1;i<len(nums);i++{
		if nums[i]==0&&nums[i-1]!=0{
			index=i
		}else{
			if index!=-1{
				if index<len(nums){
					nums[i],nums[index]=nums[index],nums[i]
				}
				for index<len(nums)&&nums[index]!=0{
					index++
				}
			}
		}
	}
}

func main() {
	
}
