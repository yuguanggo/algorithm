package main


func trap(height []int) int {
	//查出最高的柱子
	n:=len(height)
	if n==0{
		return 0
	}
	maxIndex:=0
	for i:=1;i<n;i++{
		if height[i]>height[maxIndex]{
			maxIndex=i
		}
	}
	//遍历左边的雨水
	leftSum:=0
	leftHigh:=0
	for i:=1;i<maxIndex;i++{

		if height[i]<height[leftHigh]{
			leftSum+=height[leftHigh]-height[i]
		}else{
			leftHigh=i
		}
	}
	//遍历右边的雨水
	rightSum:=0
	rightHigh:=n-1
	for i:=n-1;i>maxIndex;i--{
		if height[i]<height[rightHigh]{
			rightSum+=height[rightHigh]-height[i]
		}else{
			rightHigh=i
		}
	}
	return rightSum+leftSum
}



func main() {
	
}
