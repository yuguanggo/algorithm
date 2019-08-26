package main

import (
	"strconv"
	"fmt"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res:=make([]string,0)
	dfs(s,0,0,[]string{},&res)
	return res
}

func dfs(s string,index int,k int,path []string,res *[]string)  {
	if index>=4{
		if k==len(s){
			*res=append(*res,strings.Join(path,"."))
		}
		return
	}

	for i:=1;i<=3;i++{
		if k+i>len(s){
			break;
		}
		s1:=s[k:k+i]
		n1,_:=strconv.Atoi(s1)
		if s1[0]=='0'&&len(s1)>1||n1>255{
			continue
		}
		path=append(path,s1)
		dfs(s,index+1,k+i,path,res)
		path=path[0:len(path)-1]
	}
}

func main() {
	s:="1234";
	fmt.Println(restoreIpAddresses(s))
}
