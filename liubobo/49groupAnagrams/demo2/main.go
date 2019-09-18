package main


func groupAnagrams(strs []string) [][]string {
	record:=make(map[string][]string)
	res:=make([][]string,0)
	for i:=0;i<len(strs);i++{
		b:=[]byte(strs[i])
		qucikSort(b,0,len(b)-1)
		if _,ok:=record[string(b)];!ok{
			record[string(b)]=make([]string,0)
		}
		record[string(b)]=append(record[string(b)],strs[i])
	}
	for _,v:=range record{
		res=append(res,v)
	}
	return res
}

func qucikSort(b []byte,l,r int)  {
	if l>r{
		return
	}
	mid:=l+(r-l)/2
	v:=b[mid]
	b[mid],b[l]=b[l],b[mid]
	//三路快排
	lt:=l //b[l...lt]<v
	i:=l+1 //b[lt+1..i]==v
	gt:=r+1 //b[gt..r]>v
	for i<gt{
		if b[i]<v{
			b[i],b[lt+1]=b[lt+1],b[i]
			i++
			lt++
		}else if b[i]>v{
			b[i],b[gt-1]=b[gt-1],b[i]
			gt--
		}else{
			i++
		}
	}
	b[l],b[lt]=b[lt],b[l]
	qucikSort(b,l,lt-1)
	qucikSort(b,gt,r)
}
func main() {
	
}
