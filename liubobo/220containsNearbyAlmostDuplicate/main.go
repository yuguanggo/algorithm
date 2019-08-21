package main

import (
	"fmt"
	"math"
)

/**
220. 存在重复元素 III
给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。

示例 1:

输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:

输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func getId(x,w int) int  {
	if x<0{
		return (x+1)/w-1
	}
	return x/w
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t<0{
		return false
	}
	hash:=make(map[int]int)
	w:=t+1//桶的宽度

	for i:=0;i<len(nums);i++{
		m:=getId(nums[i],w)
		if _,ok:=hash[m];ok{
			return true
		}
		if _,ok:=hash[m-1];ok{
			if math.Abs(float64(nums[i]-hash[m-1]))<float64(w){
				return true
			}
		}
		if _,ok:=hash[m+1];ok{
			if math.Abs(float64(nums[i]-hash[m+1]))<float64(w){
				return true
			}
		}

		hash[m]=nums[i]
		if len(hash)>k{
			delete(hash,getId(nums[i-k],w))
		}
	}
	return false
}

func main() {
	nums:=[]int{1,5,9,1,5,9}
	fmt.Println(containsNearbyAlmostDuplicate(nums,2,3))
}
