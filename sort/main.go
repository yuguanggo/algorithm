package main

import (
	"algorithm/sort/mergesort"
	"algorithm/sort/sorttesthelper"
	"algorithm/sort/quicksort"
)

func main() {
	var arr7 = sorttesthelper.GenerateRandomArray(100000,1,2)
	//fmt.Println(arr)
	//insertsort.InsertSort2(arr,2,7)
	//fmt.Println(arr)
	//var arr1 = sorttesthelper.CopyIntArray(arr)
	//var arr2 = sorttesthelper.CopyIntArray(arr)
	//var arr3 = sorttesthelper.CopyIntArray(arr)
	//var arr4 = sorttesthelper.CopyIntArray(arr)
	//var arr5 = sorttesthelper.CopyIntArray(arr)
	//var arr6 = sorttesthelper.CopyIntArray(arr)
	//sorttesthelper.TestSort("Bubblesort",bubblesort.BubbleSort,arr4)
	//sorttesthelper.TestSort("SelectSort",selectsort.SelectSort,arr)
	//sorttesthelper.TestSort("SelectSort1",selectsort.SelectSort1,arr3)
	//sorttesthelper.TestSort("InsertSort",insertsort.InsertSort,arr1)
	//sorttesthelper.TestSort("ShellSort",shell_sort.ShellSort,arr5)
	//var arr7 = sorttesthelper.GenerateNearlyOrderedArray(1000000,10)
	var arr8 = sorttesthelper.CopyIntArray(arr7)
	var arr9 = sorttesthelper.CopyIntArray(arr7)
	var arr10 = sorttesthelper.CopyIntArray(arr7)
	//sorttesthelper.TestSort("InsertSort1",insertsort.InsertSort1,arr7)
	sorttesthelper.TestSort("MergeSrot",mergesort.MergeSort,arr8)
	sorttesthelper.TestSort("MergeSortBU",mergesort.MergeSortBU,arr9)
	sorttesthelper.TestSort("QuickSort",quicksort.QuickSort,arr10)
}
