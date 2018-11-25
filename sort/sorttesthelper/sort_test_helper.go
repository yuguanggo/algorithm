package sorttesthelper

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成有n个元素的随机数组,每个元素的随机范围为[rangeL, rangeR]
func GenerateRandomArray(n int, rangeL int, rangeR int) []int {
	var arr = make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = RandInt(rangeL, rangeR)
	}
	return arr
}

// 生成一个近乎有序的数组
// 首先生成一个含有[0...n-1]的完全有序数组, 之后随机交换swapTimes对数据
// swapTimes定义了数组的无序程度
func GenerateNearlyOrderedArray(n int, swapTimes int) []int {
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	for j := 1; j < swapTimes; j++ {
		m := RandInt(0, n-1)
		n := RandInt(0, n-1)
		arr[m], arr[n] = arr[n], arr[m]
	}
	return arr
}

func RandInt(min, max int) int {

	if min > max  {
		return max
	}
	return rand.Intn(max-min) + min
}

func PrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf(" %d", arr[i])
	}
	fmt.Println()
}

//判断arr数组是否有序
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func CopyIntArray(arr []int) []int {
	var carr = make([]int, len(arr))
	copy(carr, arr)
	return carr
}

func TestSort(sortname string, sort func(arr []int), arr []int) {
	start := time.Now()
	sort(arr)
	elapsed := time.Since(start).Seconds()
	if !IsSorted(arr) {
		fmt.Printf("%s:排序未完成\n", sortname)
	}
	fmt.Printf("%s:%f s\n", sortname, elapsed)
}
