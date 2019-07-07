package main

import (
	"fmt"
	"sort"
	"time"
)

/**
451. 根据字符出现频率排序
给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

示例 1:

输入:
"tree"

输出:
"eert"

解释:
'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
示例 2:

输入:
"cccaaa"

输出:
"cccaaa"

解释:
'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
注意"cacaca"是不正确的，因为相同的字母必须放在一起。
示例 3:

输入:
"Aabb"

输出:
"bbAa"

解释:
此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
注意'A'和'a'被认为是两种不同的字符

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-characters-by-frequency
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type Kv struct {
	k int
	v int
}

type KvSlice []Kv

func (kv KvSlice)Len() int  {
	return len(kv)
}

func (kv KvSlice)Less(i,j int) bool  {
	return kv[i].v>kv[j].v
}

func (kv KvSlice)Swap(i,j int)  {
	kv[i],kv[j] = kv[j],kv[i]
}


func frequencySort(s string) string {
	h:=[256]int{}
	kvs:=KvSlice{}
	for i:=0;i<len(s);i++{
		h[s[i]]++
	}
	for k,v:=range h{
		if v>0{
			kv:=Kv{
				k:k,
				v:v,
			}
			kvs = append(kvs,kv)
		}
	}
	sort.Sort(kvs)
	res:=make([]byte,0,len(s))
	for _,kv:=range kvs{
		for kv.v>0{
			res=append(res,byte(kv.k))
			kv.v--
		}
	}
	return string(res)
}


func frequencySort2(s string) string {
	idn:=256
	smap:=make([]int,256)
	idmap:=make([]int,256)
	res:=make([]byte,0,len(s))
	for _,v:=range s{
		smap[v]++
	}
	for i:=0;i<idn;i++{
		idmap[i]=i
	}
	sort.Slice(idmap, func(i, j int) bool {
		return smap[idmap[i]]>smap[idmap[j]]
	})
	for _,v:=range idmap{
		if smap[v]<=0{
			break
		}
		for smap[v]>0{
			res=append(res,byte(v))
			smap[v]--
		}
	}
	return string(res)
}

func main() {
	start:=time.Now()
	res:=""
	for i:=1;i<256000;i++{
		res=res+string(109)
	}
	elapsed:=time.Since(start).Seconds()
	fmt.Println(elapsed)
	fmt.Println(res)
	start1:=time.Now()
	res1:=make([]byte,0,256000)
	for i:=1;i<256000;i++{
		res1=append(res1,109)
	}
	fmt.Println(string(res1))
	elapsed1:=time.Since(start1).Seconds()
	fmt.Println(elapsed1)
}
