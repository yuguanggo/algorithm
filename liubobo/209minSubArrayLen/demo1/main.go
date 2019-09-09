package main

func minSubArrayLen(s int, nums []int) int {
	//窗口[l,r]最小
	l,r:=0,-1
	n:=len(nums)
	res:=n+1
	sum:=0
	for l<n{
		if r+1<n&&sum<s{
			r++
			sum+=nums[r]

		}else{
			sum-=nums[l]
			l++
		}
		if sum>=s{
			res=min(res,r-l+1)
		}
	}
	if res==n+1{
		return 0
	}
	return res
}

func min(i,j int) int  {
	if i>j{
		return j
	}
	return i
}

func minSubArrayLen2(s int, nums []int) int {
	//滑动窗口 [l..i]里和最小
	l:=0
	n:=len(nums)
	sum:=0
	res:=n+1
	for i:=0;i<n;i++{
		sum+=nums[i]
		for sum>=s{
			res=min(res,i-l+1)
			sum-=nums[l]
			l++
		}
	}
	if res==n+1{
		return 0
	}
	return res
}




func main() {
	
}
