package main

import (
	"math/rand"
	"time"
	"algorithm/heap/maxheap"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 10
	var arr []int
	var maxheap maxheap.MaxHeap
	for i := 0; i < n; i++ {
		arr = append(arr,i)
	}
	fmt.Println(arr)
	maxheap.MaxHeap(arr)
	for i := 0; i < n; i++ {
		v, err := maxheap.ExtractMax()
		if err == nil {
			fmt.Println(v)
		}
	}
	fmt.Println(maxheap)
}
