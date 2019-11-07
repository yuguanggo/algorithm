package main

/**
字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式
示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
 */

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1=="0"||num2=="0"{
		return "0"
	}
	var (
		n1  int
		n2  int
		res []int
		str string
	)
	n1 = len(num1)
	n2 = len(num2)
	res = make([]int,n1+n2-1)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			res[i+j] += int(num1[i]-'0') * int(num2[j]-'0')
			if res[i+j] >= 10 && (i+j) != 0 {
				res[i+j-1] += res[i+j] / 10
				res[i+j] = res[i+j] % 10
			}
		}
	}
	fmt.Println(res)
	for i:=0;i<len(res);i++{
		str+=strconv.Itoa(res[i])
	}
	return str
}

func main() {
  s1:="14"
  s2:="34"

  fmt.Println(multiply(s1,s2))

}
