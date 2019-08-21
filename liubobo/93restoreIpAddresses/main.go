package main

import (
	"strings"
	"strconv"
)

/**
93. 复原IP地址
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func restoreIpAddresses(s string) []string {
	var(
		res []string
	)
	backTracking(s,0,[]string{},&res)
	return res
}

func backTracking(s string,pos int,path []string,res *[]string)  {
	if len(path)>=4{
		if pos==len(s){
			*res=append(*res,strings.Join(path,"."))
		}
		return
	}
	for i:=1;i<=3;i++{
		if pos+i>len(s){
			break
		}
		subs:=s[pos:pos+i]
		n,_:=strconv.Atoi(subs);
		if strings.HasPrefix(subs,"0")&&len(subs)>1||n>255{
			continue
		}
		path=append(path,subs)
		backTracking(s,pos+i,path,res)
		path=path[0:len(path)-1]
	}
}

func main() {
	
}
