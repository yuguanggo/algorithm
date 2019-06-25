package main

import "fmt"

func partition1(nums []int, l int, r int) int {
	//随机找个点 防止退化成on2
	mid := (l + r) / 2
	nums[l], nums[mid] = nums[mid], nums[l]
	v := nums[l]
	j := l //j一直指向小于v的末尾
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

//单路快排
func quickSort1(nums []int, l int, r int) {
	var p int
	if l>r{
		return
	}
	p = partition1(nums, l, r)
	quickSort1(nums, l, p-1)
	quickSort1(nums, p+1, r)
}

func partition2(nums []int,l int,r int) int  {
	mid := (l+r)/2
	nums[l],nums[mid] = nums[mid],nums[l]
	v:=nums[l]
	i:=l+1
	j:=r
	for{
		//不用小于等于v是因为数组有大量重复值时，两端会分配不均
		for i<=r&&nums[i]<v{
			i++
		}
		for j>=l+1&&nums[j]>v{
			j--
		}
		if i>j{
			break
		}
		nums[i],nums[j] =nums[j],nums[i]
		i++
		j--
	}
	nums[l],nums[j] = nums[j],nums[l]
	return j
}

//双路快排
func quickSort2(nums []int,l int,r int)  {
	var p int
	if l>r{
		return
	}
	p = partition2(nums,l,r)
	quickSort2(nums,l,p-1)
	quickSort2(nums,p+1,r)
}

func quickSort3(nums []int,l int,r int)  {
	if l>=r {
		return
	}
	mid:=(l+r)/2
	nums[l],nums[mid] = nums[mid],nums[l]
	lt:=l//最小位的末尾
	gt:=r+1//最大位的首位
	i:=l+1
	v:=nums[l]
	for i<gt{
		if nums[i]<v{
			nums[i],nums[lt+1] = nums[lt+1],nums[i]
			i++
			lt++
		}else if nums[i]>v{
			nums[i],nums[gt-1] = nums[gt-1],nums[i]
			gt--
		}else{
			i++
		}
	}
	nums[l],nums[lt] =nums[lt],nums[l]
	quickSort3(nums,l,lt-1)
	quickSort3(nums,gt,r)
}


func main() {
	var arr = []int{44, 2, 32, 42, 3, 5, 321, 76, 3, 22,90}
	quickSort3(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
