package main

import (
	"fmt"
	"strconv"
)

/**
第k个排列
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
i是下表
X = ai*(n-1)!+a(i-1)*(n-2)!....a1*0!

逆康托展开举例
一开始已经提过了，康托展开是一个全排列到一个自然数的双射，因此是可逆的。即对于上述例子，在  给出61可以算出起排列组合为34152。由上述的计算过程可以容易的逆推回来，具体过程如下：
用 61 / 4! = 2余13，说明  ,说明比首位小的数有2个，所以首位为3。
用 13 / 3! = 2余1，说明  ，说明在第二位之后小于第二位的数有2个，所以第二位为4。
用 1 / 2! = 0余1，说明  ，说明在第三位之后没有小于第三位的数，所以第三位为1。
用 1 / 1! = 1余0，说明  ，说明在第二位之后小于第四位的数有1个，所以第四位为5。
最后一位自然就是剩下的数2。
通过以上分析，所求排列组合为 34152。

3/2! = 1余1 2
1/1! = 1余0 3
0/1  = 0 1
再举个例子说明。
　　在  5个数的排列组合中，计算 34152的康托展开值。
首位是3，则小于3的数有两个，为1和2，  ，则首位小于3的所有排列组合为2(5-1)!
第二位是4，由于第一位小于4，1、2、3中一定会有1个充当第一位，所以排在4之下的只剩2个，所以其实计算的是在第二位之后小于4的个数。因此。2(4-1)!
第三位是1，则在其之后小于1的数有0个，所以  。0(3-1)!
第四位是5，则在其之后小于5的数有1个，为2，所以  。2(2-1)!
最后一位就不用计算啦，因为在它之后已经没有数了，所以  固定为0
根据公式：

　　所以比34152小的组合有61个，即34152是排第62。
 */
func getPermutation(n int, k int) string {
	var arr []int
	var res string
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	k--
	for j := n; j >= 1; j-- {
		m := k / factorial(j-1)
		k = k % factorial(j-1)
		res += strconv.Itoa(arr[m])
		arr = append(arr[:m], arr[(m + 1):]...)
	}
	return res
}



func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println(getPermutation(3,3))
	str:="1232"
	fmt.Println(str[0]-'0')
}
