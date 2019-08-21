package quicksort

//对[l...r]区间进行排序
//是的[l+1..j]<v  [j+1...i-1]>v i是分割点
//v|---<v-----|--->v----|e--
//l|---------j|---------|i
func partition(nums []int, l, r int) int {
	if l>=r{
		return l
	}
	//选择中间的点，防止近乎有序的数组算法退化成o(n2)
	mid := l + (r-l)/2
	nums[l], nums[mid] = nums[mid], nums[l]
	v := nums[l]
	j:= l//指向小于v的末尾
	for i:=l+1; i <=r; i++ {
		if nums[i] < v {
			j++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[j],nums[l]=nums[l],nums[j]
	return j
}

//单路快排
func QuickSort(nums []int, l, r int) {
	//[l...r]的区间进行排序
	if l>=r{
		return
	}
	var p = partition(nums, l, r)
	QuickSort(nums,l,p-1)
	QuickSort(nums,p+1,r)
}

//a[l+1..i]<v a[j..r]>v
func partition2(nums []int,l,r int) int  {
	mid:=l+(r-l)/2
	nums[l],nums[mid]=nums[mid],nums[l]
	v:=nums[l]
	//arr[l+1..i)<=v arr(j..r]>=v
	i:=l+1
	j:=r
	for {
		//不用num[i]<=v 是为了防止数组元素都相同的时候造成树不平衡
		for i<=r&&nums[i]<v{
			i++
		}
		for j>=l+1&&nums[j]>v{
			j--
		}
		if i>j{
			break
		}
		nums[i],nums[j] = nums[j],nums[i]
		i++
		j--
	}
	nums[l],nums[i]=nums[i],nums[l]
	return i
}

//双路快排
func QuickSort2(nums []int,l,r int)  {
	if l>=r{
		return
	}
	var p= partition2(nums,l,r)
	QuickSort2(nums,l,p-1)
	QuickSort2(nums,p+1,r)
}

//nums[l..lt]<v nums[lt+1...i-1]=v nums[gt..r]>v
func QuickSort3(nums []int,l,r int)  {
	if l>=r{
		return
	}
	mid:=l+(r-l)/2
	nums[l],nums[mid] = nums[mid],nums[l]
	v:=nums[l]
	lt:=l
	gt:=r+1
	i:=l+1
	for i<gt{
		if nums[i]>v{
			nums[gt-1],nums[i]=nums[i],nums[gt-1]
			gt--
		}else if nums[i]<v{
			nums[lt+1],nums[i]=nums[i],nums[lt+1]
			i++
			lt++
		}else{
			i++
		}
	}
	nums[l],nums[lt] = nums[lt],nums[l]
	QuickSort3(nums,l,lt-1)
	QuickSort3(nums,gt,r)
}