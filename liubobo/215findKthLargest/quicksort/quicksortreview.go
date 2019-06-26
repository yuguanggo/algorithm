package quicksort

//单路快排
//arr[l+1...j]<v  arr[j+1..r]>v

func repartition(nums []int, l,r int) int  {
	mid:=l+(r-1)/2
	nums[l],nums[mid]=nums[mid],nums[l]
	v:=nums[l]
	j:=l//
	for i:=l+1;i<=r;i++{
		if nums[i]<v{
			j++
			nums[i],nums[j]=nums[j],nums[i]
		}
	}
	nums[j],nums[l]=nums[l],nums[j]
	return j
}

func ReSort(nums []int,l,r int)  {
	if l>=r{
		return
	}
	var p = repartition(nums,l,r)
	ReSort(nums,l,p-1)
	ReSort(nums,p+1,r)
}

//双路快排
//arr[l+1...i-1]<v arr[j+1...r]>v
func repartition2(nums []int,l,r int) int  {
	mid:=l+(r-l)/2
	nums[mid],nums[l]=nums[l],nums[mid]
	v:=nums[l]
	i:=l+1
	j:=r
	for{
		for nums[i]<v&&i<=r{
			i++
		}
		for nums[j]>v&&j>=l+1{
			j--
		}
		if i>j{
			break
		}
		nums[i],nums[j]=nums[j],nums[i]
		i++
		j--
	}
	nums[l],nums[j]=nums[j],nums[l]
	return j
}


func ReSort2(nums []int,l,r int)  {
	if l>=r{
		return
	}
	var p=repartition2(nums,l,r)
	ReSort2(nums,l,p-1)
	ReSort2(nums,p+1,r)
}

//nums[l+1..lt]<v nums[lt+1..i]==v nums[gt..r]>v
func ReSort3(nums []int,l,r int)  {
	if l>=r{
		return
	}
	mid:=l+(r-l)/2
	nums[mid],nums[l]=nums[l],nums[mid]
	v:=nums[l]
	lt:=l
	gt:=r+1
	i:=l+1
	for i<gt{
		if nums[i]>v{
			nums[i],nums[gt-1]=nums[gt-1],nums[i]
			gt--
		}else if nums[i]<v{
			nums[i],nums[lt+1] =nums[lt+1],nums[i]
			i++
			lt++
		}else{
			i++
		}
	}
	nums[l],nums[lt]=nums[lt],nums[l]
	ReSort3(nums,l,lt-1)
	ReSort3(nums,gt,r)
}