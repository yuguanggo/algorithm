package main

import "fmt"

func trap(height []int) int {
	n:=len(height)
	if n==0{
		return 0
	}
	maxH:=height[0]
	maxIndex:=0

	for i:=1;i<n;i++{
		if height[i]>maxH{
			maxH=height[i]
			maxIndex=i
		}
	}
	//从左遍历
	lres:=0
	lmax:=height[0]
	for i:=1;i<maxIndex;i++{
		if height[i]<lmax{
			lres+=lmax-height[i]
		}else{
			lmax=height[i]
		}
	}
	rres:=0
	rmax:=height[n-1]
	for i:=n-1;i>maxIndex;i--{
		if height[i]<rmax{
			rres+=rmax-height[i]
		}else{
			rmax=height[i]
		}
	}
	return lres+rres
}




func main() {
	var height = []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))
}
