package main

import "math"

/**
4. 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n:=len(nums1)
	m:=len(nums2)
	if(n>m){
		return findMedianSortedArrays(nums2,nums1)
	}
	var lmax1,rmin1,lmax2,rmin2 int
	low:=0
	high:=2*n
	for low<=high{
		c1:=(low+high)/2
		c2:=m+n-c1
		if c1==0{
			lmax1=math.MinInt32
		}else{
			lmax1=nums1[(c1-1)/2]
		}
		if c1==2*n{
			rmin1=math.MaxInt32
		}else{
			rmin1=nums1[c1/2]
		}
		if c2==0{
			lmax2=math.MinInt32
		}else{
			lmax2=nums2[(c2-1)/2]
		}
		if c2==2*m{
			rmin2=math.MaxInt32
		}else{
			rmin2=nums2[c2/2]
		}
		if lmax1>rmin2{
			high=c1-1
		}else if lmax2>rmin1{
			low=c1+1
		}else{
			break
		}
	}
	return float64(max(lmax1,lmax2)+min(rmin1,rmin2))/2
}
func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func min(i,j int) int  {
	if i<j{
		return i
	}
	return j
}


func main() {
	
}
