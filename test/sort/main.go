package main

import "fmt"

//稳定排序
func BubbleSort(arr []int)  {
	n:=len(arr)
	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
}

func SelectSort(arr []int)  {
	n:=len(arr)
	for i:=0;i<n;i++{
		minIndex := i
		for j:=i+1;j<n;j++{
			if arr[j]<arr[minIndex]{
				minIndex=j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}

func SelectSort2(arr []int)  {
	n:=len(arr)
	left:=0;
	right:=n-1;
	for left<right{
		minIndex:=left
		maxIndex:=right
		if arr[minIndex]>arr[maxIndex]{
			arr[minIndex],arr[maxIndex]=arr[maxIndex],arr[minIndex]
		}
		for i:=left+1;i<right;i++{
			if arr[i]<arr[minIndex]{
				minIndex=i
			}
			if arr[i]>arr[maxIndex]{
				maxIndex=i
			}
		}
		arr[left],arr[minIndex]=arr[minIndex],arr[left]
		arr[right],arr[maxIndex]=arr[maxIndex],arr[right]
		left++
		right--
	}
}

func main() {
	arr:=[]int{5,4,3,2,1}
	SelectSort2(arr)
	fmt.Println(arr)
}
