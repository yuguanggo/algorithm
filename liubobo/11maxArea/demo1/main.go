package main

func maxArea(height []int) int {
	//双指针法 移动短板 不会导致错失最大面积
	l:=0
	r:=len(height)-1
	res:=0
	for l<r{
		if height[l]>height[r]{
			res = max(res,height[r]*(r-l))
			r--
		}else{
			res = max(res,height[l]*(r-l))
			l++
		}
	}
	return res
}

func max(i,j int ) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
