package main

import (
	"fmt"
	"math"
)

func main() {
	s := []float64{0.1,0.0,3.0,-2.0,0.9,-1.3,5.0,-4.4}
	fmt.Println(maxProduct(s))
}
/*
 * 子数组最大乘积
 */
func maxProduct( arr []float64 ) float64 {
	// write code here
	res := arr[0]
	max := arr[0]
	min := arr[0]
	for _,v := range arr[1:]{
		max,min = math.Max(math.Max(v*max,v),v*min),math.Min(math.Min(v*min, v),v*max) //当前元素结尾的最大&&小值
		if max > res{
			res = max
		}
	}
	return res
}