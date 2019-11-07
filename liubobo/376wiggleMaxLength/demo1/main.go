package main


func wiggleMaxLength2(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	up:=make([]int,n)
	down:=make([]int,n)
	up[0]=1
	down[0]=1
	for i:=1;i<n;i++{
		if nums[i]>nums[i-1]{
			up[i]=down[i-1]+1
			down[i]=down[i-1]
		}else if nums[i]<nums[i-1]{
			down[i]=up[i-1]+1
			up[i]=up[i-1]
		}else{
			up[i]=up[i-1]
			down[i]=down[i-1]
		}
	}
	if up[n-1]>down[n-1]{
		return up[n-1]
	}
	return down[n-1]
}

func wiggleMaxLength(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	up:=1
	down:=1
	for i:=1;i<n;i++{
		if nums[i]>nums[i-1]{
			up=down+1
		}else if nums[i]<nums[i-1]{
			down=up+1
		}
	}
	if up>down{
		return up
	}
	return down
}

func main() {
	
}
