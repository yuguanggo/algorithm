package main

import (
	"strconv"
	"fmt"
)

/**
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
 */
func restoreIpAddresses(s string) []string {
	var(
		res []string
	)
	findIp(s,0,0,"",&res)
	return res
}

func findIp(s string,f int,idx int,ip string,res *[]string)  {
	if(idx==3){
		if len(s)-f<=3{
			if s[f]=='0'&&(f!=len(s)-1){
				return
			}
			num,_:=strconv.Atoi(s[f:len(s)])
			if num<256{
				*res = append(*res,ip+strconv.Itoa(num))
			}
		}
	}else{
		for i:=1;i<=3;i++{
			if f+i>=len(s){
				break
			}
			num,_ := strconv.Atoi(s[f:f+i])
			if num<256{
				findIp(s,f+i,idx+1,ip+strconv.Itoa(num)+".",res)
			}
			if s[f]=='0'&&i==1{
				break
			}

		}
	}
}


func main() {
	s:="010010"
	fmt.Println(restoreIpAddresses(s))
}
