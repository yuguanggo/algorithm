package main

import (
	"algorithm/search/binarysearch"
	"math/rand"
	"fmt"
	"time"
)

func main() {
	//var arr = sorttesthelper.GenerateNearlyOrderedArray(100000, 0)
	//index := binarysearch.Find(arr, 1)
	//fmt.Println(index)
	//index2 := binarysearch.RecursionFind(arr, 0,len(arr)-1,1)
	//fmt.Println(index2)
	var (
		bts *binarysearch.BTS
		i int
		rannun int
	)
	bts = binarysearch.Init()
	rand.Seed(time.Now().UnixNano())
	for i = 0;i<10;i++{
		rannun = rand.Intn(100)
		fmt.Println(rannun)
		bts.Insert(i,rannun)
	}
	fmt.Println(bts.GetRoot())

}
