package main

func sort(nums []int,l,r int)  {
	if l>r{
		return
	}
	mid:=l+(r-l)/2
	nums[mid],nums[l]=nums[l],nums[mid]
	v:=nums[l]
	lt:=l//[l...lt]<v
	gt:=r+1 //[gt..r]>v
	i:=l+1 //[li+1..i-1]==v
	for i<gt{
		if nums[i]<v{
			nums[i],nums[lt+1]=nums[lt+1],nums[i]
			i++
			lt++
		}else if nums[i]>v{
			nums[i],nums[gt-1]=nums[gt-1],nums[i]
			gt--
		}else{
			i++
		}
	}
	nums[l],nums[lt]=nums[lt],nums[l]
	sort(nums,l,lt-1)
	sort(nums,gt,r)
}
func main() {
	
}
