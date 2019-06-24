package main

import "fmt"

/**
 数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
方法一 : 任意一种O(nlogn)算法对数组进行降序排序, 取下标为k - 1的数组元素即可
方法二 : 使用堆, 维护一个元素个数为k的最大堆, 将所有数字放入到最大堆中, 全部放完之后, 最大堆中最小的那个元素就是第k个最大的元素
方法三 : 快速排序思想, 使用三路快排, 每次都将数组分割成三部分, 每次只需要在其中一部分继续寻找, 时间复杂度O(logn), 方法三代码如下:
---------------------
作者：七夜丶雪
来源：CSDN
原文：https://blog.csdn.net/love905661433/article/details/84930799
版权声明：本文为博主原创文章，转载请附上博文链接！
 */

func findKthLargest(nums []int, k int) int {
	return quickSort(nums,0,len(nums)-1,k)
}

func quickSort(nums []int,l int,r int,k int) int  {
	mid:=(l+r)/2
	nums[l],nums[mid]=nums[mid],nums[l]
	v:=nums[l]
	lt:=l
	gt:=r+1
	i:=l+1
	for i<gt{
		if nums[i]>v{
			nums[i],nums[lt+1] = nums[lt+1],nums[i]
			lt++
			i++
		}else if nums[i]<v{
			nums[i],nums[gt-1] = nums[gt-1],nums[i]
			gt--
		}else {
			i++
		}
	}
	nums[l],nums[lt] = nums[lt],nums[l]
	if (k-1)<=lt-1{
		return quickSort(nums,l,lt-1,k)
	}else if (k-1)>=gt{
		return quickSort(nums,gt,r,k)
	}else {
		return nums[k-1]
	}
}

func quickSort2(nums []int,l int,r int)   {
	if l>=r{
		return
	}
	mid:=(l+r)/2
	nums[l],nums[mid]=nums[mid],nums[l]
	v:=nums[l]
	lt:=l
	gt:=r+1
	i:=l+1
	for i<gt{
		if nums[i]<v{
			nums[i],nums[lt+1] = nums[lt+1],nums[i]
			lt++
			i++
		}else if nums[i]>v{
			nums[i],nums[gt-1] = nums[gt-1],nums[i]
			gt--
		}else {
			i++
		}
	}
	nums[l],nums[lt] = nums[lt],nums[l]
	quickSort2(nums,l,lt)
	quickSort2(nums,gt,r)

}


func main() {
	var nums =[]int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(nums,2))
	//quickSort2(nums,0,(len(nums)-1))
	//fmt.Println(nums)

}
