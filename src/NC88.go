package main

import "fmt"

func main() {
	a := []int{-1,2,0}
	fmt.Println(findKth(a,6,1))
}
/*
有一个整数数组，请你根据快速排序的思路，找出数组中第K大的数。

给定一个整数数组a,同时给定它的大小n和要找的K(K在1到n之间)，请返回第K大的数，保证答案存在
 */

/*
快排就行
 */
func findKth( a []int ,  n int ,  K int ) int {
	// write code here
	for ;;{
		rt := []int{}
		lf := []int{}
		for _,val := range a[1:]{
			if val >= a[0]{
				rt = append(rt, val) //大的放rt
			}else{
				lf = append(lf, val)
			}
		}
		if len(rt)+1==K{
			return a[0]
		}else if len(rt)+1>K{
			a = rt
			//fmt.Println(K,a)
		}else{
			K -= len(rt)+1
			a = lf
			//fmt.Println(K,a)
		}
	}
}