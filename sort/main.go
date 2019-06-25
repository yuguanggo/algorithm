package main

import (
	"algorithm/sort/sorttesthelper"
	"algorithm/sort/bubblesort"
	"algorithm/sort/selectsort"
)

func main() {
	var arr = sorttesthelper.GenerateRandomArray(100000,1,2)
	//fmt.Println(arr)
	//insertsort.InsertSort2(arr,2,7)
	//fmt.Println(arr)
	//var arr1 = sorttesthelper.CopyIntArray(arr)
	//var arr2 = sorttesthelper.CopyIntArray(arr)
	var arr3 = sorttesthelper.CopyIntArray(arr)
	var arr4 = sorttesthelper.CopyIntArray(arr)
	//var arr5 = sorttesthelper.CopyIntArray(arr)
	//var arr6 = sorttesthelper.CopyIntArray(arr)
	sorttesthelper.TestSort("Bubblesort",bubblesort.BubbleSort,arr4)
	sorttesthelper.TestSort("SelectSort",selectsort.SelectSort,arr)
	sorttesthelper.TestSort("SelectSort1",selectsort.SelectSort1,arr3)
	//sorttesthelper.TestSort("InsertSort",insertsort.InsertSort,arr1)
	//sorttesthelper.TestSort("ShellSort",shell_sort.ShellSort,arr5)
<<<<<<< HEAD
	var arr7 = sorttesthelper.GenerateNearlyOrderedArray(1000000,300)
	var arr8 = sorttesthelper.CopyIntArray(arr7)
	sorttesthelper.TestSort("InsertSort1",insertsort.InsertSort1,arr7)
	sorttesthelper.TestSort("MergeSrot",mergesort.MergeSrot,arr8)
	var a []int
=======
	//var arr7 = sorttesthelper.GenerateNearlyOrderedArray(100000,100000)
	//var arr8 = sorttesthelper.CopyIntArray(arr7)
	//var arr9 = sorttesthelper.CopyIntArray(arr7)
	//var arr10 = sorttesthelper.CopyIntArray(arr7)
	//var arr11 = sorttesthelper.CopyIntArray(arr7)
	//var arr12 = sorttesthelper.CopyIntArray(arr7)
	//var arr13 = sorttesthelper.CopyIntArray(arr7)
	//var arr14 = sorttesthelper.CopyIntArray(arr7)
	//var arr15 = sorttesthelper.CopyIntArray(arr7)
	////sorttesthelper.TestSort("InsertSort1",insertsort.InsertSort1,arr7)
	//sorttesthelper.TestSort("MergeSrot",mergesort.MergeSort,arr8)
	//sorttesthelper.TestSort("MergeSortBU",mergesort.MergeSortBU,arr9)
	//sorttesthelper.TestSort("QuickSort",quicksort.QuickSort,arr10)
	//sorttesthelper.TestSort("QuickSort2",quicksort.QuickSort2,arr11)
	//sorttesthelper.TestSort("QuickSort3",quicksort.QuickSort3,arr12)
	//sorttesthelper.TestSort("HeapSort",maxheap.HeapSort,arr13)
	//sorttesthelper.TestSort("HeapSort2",maxheap.HeapSort2,arr14)
	//sorttesthelper.TestSort("HeapSort3",maxheap.HeapSort3,arr15)
>>>>>>> b0b2b96777a3c89a52a083eb2703da3db5bd6ea7
}
