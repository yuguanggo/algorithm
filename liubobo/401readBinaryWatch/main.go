package main

import (
	"strconv"
	"sort"
)

/**
401. 二进制手表
二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。

每个 LED 代表一个 0 或 1，最低位在右侧。
例如，上面的二进制手表读取 “3:25”。

给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

案例:

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
 

注意事项:

输出的顺序没有要求。
小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-watch
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func readBinaryWatch(num int) []string {
	var watch=[]int{1,2,4,8,1,2,4,8,16,32}
	var res=[]string{}
	helper(watch,num,0,[2]int{0,0},&res)
	sort.Strings(res)
	return  res
}
func helper(watch []int,num int,index int,path [2]int,res *[]string)  {
	if num<0{
		return
	}
	if num==0{
		if path[0]<=11&&path[1]<=59{
			s:=format(path)
			*res=append(*res,s)
		}
		return
	}
	for i:=index;i<len(watch);i++{
		if i<=3{
			//小时
			path[0]=path[0]+watch[i]
		}else{
			path[1]=path[1]+watch[i]
		}
		helper(watch,num-1,i+1,path,res)
		if i<=3{
			//小时
			path[0]=path[0]-watch[i]
		}else{
			path[1]=path[1]-watch[i]
		}
	}
}

func format(path [2]int) string  {

	var h,i string
	h=strconv.Itoa(path[0])
	if path[1]<10{
		i="0"+strconv.Itoa(path[1])
	}else{
		i=strconv.Itoa(path[1])
	}
	return h+":"+i
}

func main() {
	
}
