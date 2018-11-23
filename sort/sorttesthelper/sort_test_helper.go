package sorttesthelper

import (
	"math/rand"
	"time"
	"fmt"
)

// 生成有n个元素的随机数组,每个元素的随机范围为[rangeL, rangeR]
func GenerateRandomArray(n int, rangeL int, rangeR int) []int {
	var arr = make([]int,n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = RandInt(rangeL,rangeR)
	}
	return arr
}

func RandInt(min, max int) int {
	if min > max || min == 0 || max == 0{
		return max
	}
	return rand.Intn(max-min)+min
}

func PrintArray(arr []int)  {
	for i:=0;i<len(arr) ;i++  {
		fmt.Printf(" %d",arr[i])
	}
	fmt.Println()
}

//判断arr数组是否有序
func IsSorted(arr []int) bool{
	for i:=0;i<len(arr)-1 ;i++  {
		if arr[i]>arr[i+1]{
			return false
		}
	}
	return true
}

func CopyIntArray(arr []int) []int {
	var carr = make([]int,len(arr))
	copy(carr,arr)
	return carr
}

func TestSort(sortname string,sort func(arr []int,n int),arr []int,n int)  {
	start:=time.Now()
	sort(arr,n)
	elapsed:=time.Since(start)
	if !IsSorted(arr){
		fmt.Printf("%s:排序未完成\n",sortname)
	}
	fmt.Printf("%s:%s\n",sortname,elapsed)
}
