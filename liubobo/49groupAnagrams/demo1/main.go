package main

import (
	"fmt"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	res:=make([][]string,0)
	record:=make(map[string][]string)
	for i:=0;i<len(strs);i++{
		count:=[26]int{}
		for j:=0;j<len(strs[i]);j++{
			count[strs[i][j]-'a']++
		}
		s:=""
		for k:=0;k<26;k++{
			s+="#"
			s+=strconv.Itoa(count[k])
		}
		if _,ok:=record[s];!ok{
			record[s]=make([]string,0)

		}
		record[s]=append(record[s],strs[i])
	}
	for _,v:=range record{
		res=append(res,v)
	}
	return res
}


func main() {
	s:=[]string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(s))
}
