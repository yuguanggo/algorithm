package main

import "fmt"

func bs(numbers []int,l,r ,target int) int  {

	for l<=r{
		mid:=l+(r-l)/2
		if numbers[mid]==target{
			return mid
		}else if numbers[mid]>target{
			r=mid-1
		}else{
			l=mid+1
		}
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	n:=len(numbers)
	for i:=0;i<n;i++{
		p:=bs(numbers,i+1,n-1,target-numbers[i])
		if p!=-1{
			return []int{i+1,p+1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{-1,0},-1))
}
