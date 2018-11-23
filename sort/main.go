package main

import (
	"algorithm/sort/sorttesthelper"
	"algorithm/sort/insertsort"
	"algorithm/sort/selectsort"
)

func main() {
	var arr = sorttesthelper.GenerateRandomArray(10000,1,10000)
	var arr1 = sorttesthelper.CopyIntArray(arr)
	var arr2 = sorttesthelper.CopyIntArray(arr)
	sorttesthelper.TestSort("SelectSort",selectsort.SelectSort,arr,len(arr))
	sorttesthelper.TestSort("InsertSort",insertsort.InsertSort,arr1,len(arr1))
	sorttesthelper.TestSort("InsertSort1",insertsort.InsertSort1,arr2,len(arr2))
	//sorttesthelper.PrintArray(arr)
}
