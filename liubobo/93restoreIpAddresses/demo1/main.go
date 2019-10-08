package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res:=make([]string,0)
	helper(s,0,[]string{},&res)
	return res
}

func helper(s string,pos int,path []string,res *[]string)  {
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
		sub:=s[pos:pos+i]
		num,_:=strconv.Atoi(sub)
		if sub[0]=='0'&&len(sub)>1||num>255{
			continue
		}
		path=append(path,sub)
		helper(s,pos+i,path,res)
		path=path[:len(path)-1]
	}
}
func main() {
	
}
