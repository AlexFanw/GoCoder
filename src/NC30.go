package main

import "fmt"

func main() {
	s := []int{4,3,2,1,-1}
	fmt.Println(minNumberdisappered(s))
}

//func minNumberdisappered( arr []int ) int {
//	// write code here
//	minNum := 1
//	for _,val := range arr{
//		if val<=0{
//			continue
//		}
//		if val <= minNum{
//			minNum++
//		}
//	}
//	return minNum
//}

func minNumberdisappered( arr []int) int{
	min := 0
	max := 0
	sum := 0
	for _,val := range arr{
		if val < min{
			min = val
		}
		if val > max{
			max = val
		}
		sum += val
	}
	if sum == (max+min)*(max-min+1)/2{
		return max+1
	}else {
		return (max+min)*(max-min+1)/2 - sum
	}
}