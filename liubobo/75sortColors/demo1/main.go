package main

func sortColors(nums []int)  {
	n:=len(nums)
	l:=-1 //[0..l]==1
	r:=n //[r..n-1]==3
	i:=0 //[l+1..i]==2
	for i<r{
		if nums[i]==0{
			nums[i],nums[l+1]=nums[l+1],nums[i]
			i++
			l++
		}else if nums[i]==1{
			i++
		}else{
			nums[i],nums[r-1]=nums[r-1],nums[i]
			r--
		}
	}

}

func quickSort(nums []int,l,r int)  {
	if l>r{
		return
	}
	//三路快排 [l,lt]<v [lt+1..i]==v [gt..r]>v
	mid:=l+(r-l)/2
	nums[l],nums[mid]=nums[mid],nums[l]
	v:=nums[l]
	lt:=l
	i:=l+1
	gt:=r+1
	for i<gt{
		if nums[i]>v{
			nums[i],nums[gt-1]=nums[gt-1],nums[i]
			gt--
		}else if nums[i]<v{
			nums[i],nums[lt+1]=nums[lt+1],nums[i]
			lt++
			i++
		}else{
			i++
		}
	}
	nums[l],nums[lt]=nums[lt],nums[l]
	quickSort(nums,l,lt-1)
	quickSort(nums,gt,r)
}

func main() {
	
}
