package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	record:=make(map[int]int)
	res:=0
	for i:=0;i<len(A);i++{
		for j:=0;j<len(B);j++{
			record[A[i]+B[j]]++
		}
	}
	for i:=0;i<len(C);i++{
		for j:=0;j<len(D);j++{
			if record[0-C[i]-D[j]]!=0{
				res+=record[0-C[i]-D[j]]
			}
		}
	}
	return res
}


func main() {
	
}
