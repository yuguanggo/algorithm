package main

import (
	"math/rand"
	"algorithm/heap/maxheap"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//n := 10
	//var arr []int
	//var maxheap maxheap.MaxHeap
	//for i := 0; i < n; i++ {
	//	arr = append(arr,i)
	//}
	//fmt.Println(arr)
	//maxheap.MaxHeap(arr)
	//for i := 0; i < n; i++ {
	//	v, err := maxheap.ExtractMax()
	//	if err == nil {
	//		fmt.Println(v)
	//	}
	//}
	//fmt.Println(maxheap)
	var (
		n int = 10
		v int
		indexMaxHeap *maxheap.IndexMaxHeap
	)
	indexMaxHeap = maxheap.Init(n)
	for i:=0;i<n;i++{
		v = rand.Intn(100)
		indexMaxHeap.Insert(i,v)
	}
	fmt.Println(indexMaxHeap)
	//max := indexMaxHeap.ExtractMax()
	//fmt.Println(max)
	indexMaxHeap.Change(5,200)
	fmt.Println(indexMaxHeap)
}
