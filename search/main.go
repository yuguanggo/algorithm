package main

import (
	"algorithm/sort/sorttesthelper"
	"fmt"
	"algorithm/search/binarysearch"
)

func main() {
	var arr = sorttesthelper.GenerateNearlyOrderedArray(100000, 0)
	index := binarysearch.Find(arr, 1)
	fmt.Println(index)
	index2 := binarysearch.RecursionFind(arr, 0,len(arr)-1,1)
	fmt.Println(index2)
}
