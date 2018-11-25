package mergesort

import "algorithm/sort/insertsort"

func merge(arr []int,left int,middle int,right int)  {
	var size =right-left+1
	var cparr = make([]int,size)
	for m:=left;m<=right;m++{
		cparr[m-left] = arr[m]
	}
	//i指向左边的数组起始位置，j指向右边的数据起始位置
	i,j:=left,middle+1
	for k:=left;k<=right;k++{
		if i>middle{
			//左边边已经排序完
			arr[k] = cparr[j-left]
			j++
		}else if j>right{
			arr[k] = cparr[i-left]
			i++
		}else if cparr[i-left]<cparr[j-left]{
			arr[k] = cparr[i-left]
			i++
		}else{
			arr[k] = cparr[j-left]
			j++
		}
	}
}

func sort(arr []int,left int,right int)  {
	//if left>=right{
	//	return
	//}
	if right-left<=15 {
		insertsort.InsertSort2(arr,left,right)
		return
	}
	middle:=(left+right)/2
	sort(arr,left,middle)
	sort(arr,middle+1,right)
	// 对于arr[mid] <= arr[mid+1]的情况,不进行merge
	// 对于近乎有序的数组非常有效,但是对于一般情况,有一定的性能损失
	if arr[middle]>arr[middle+1]{
		merge(arr,left,middle,right)
	}

}

func MergeSrot(arr []int)  {
	sort(arr,0,len(arr)-1)
}
