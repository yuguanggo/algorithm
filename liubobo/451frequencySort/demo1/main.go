package main

import "sort"

func frequencySort(s string) string {
	sm:=make([]int,256)
	idm:=make([]int,256)
	res:=make([]byte,0)
	for i:=0;i<len(s);i++{
		sm[s[i]]++
	}
	for i:=0;i<256;i++{
		idm[i]=i
	}
	sort.Slice(idm, func(i, j int) bool {
		return sm[idm[i]]>sm[idm[j]]
	})
	for i:=0;i<256;i++{
		if sm[idm[i]]<=0{
			break
		}
		for sm[idm[i]]>0{
			res=append(res,byte(idm[i]))
			sm[idm[i]]--
		}
	}
	return string(res)
}


func main() {
	
}
