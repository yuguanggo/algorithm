package main

import "math"

func maxProfit(prices []int) int {
	minPrice:=math.MaxInt32
	maxPorfit:=0
	for i:=0;i<len(prices);i++{
		if prices[i]<minPrice{
			minPrice=prices[i]
		}else{
			if prices[i]-minPrice>maxPorfit{
				maxPorfit=prices[i]-minPrice
			}
		}
	}
	return maxPorfit
}


func main() {
	
}
