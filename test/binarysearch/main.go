package main

import (
	"algorithm/sort/sorttesthelper"
	"fmt"
)

func binarySearch(arr []int,target int) int  {
	//非递归 在[l....r]区间查找target
	l:=0
	r:=len(arr)-1
	for l<=r{
		mid:=l+(r-l)/2 //防止整型溢出
		if arr[mid]==target{
			return mid
		}
		if arr[mid]>target{
			r=mid-1
		}else{
			l=mid+1
		}
	}
	return -1
}

func recursionBinarySearch(arr []int,l int,r int,target int) int  {
	//在[l....r]的区间查找target
	if l>r{
		return -1
	}
	mid:=l+(r-l)/2
	if arr[mid]==target{
		return mid
	}
	if arr[mid]>target{
		return recursionBinarySearch(arr,l,mid-1,target)
	}else{
		return recursionBinarySearch(arr,mid+1,r,target)
	}
}

func main() {
	 dd:=sorttesthelper.GenerateNearlyOrderedArray(10000,0)
	 fmt.Println(binarySearch(dd,2))
	 fmt.Println(recursionBinarySearch(dd,0,len(dd)-1,3))
}
