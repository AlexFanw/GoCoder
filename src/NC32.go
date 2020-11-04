package main

import "fmt"

func main() {
	fmt.Println(sqrt(10))
}

func sqrt( x int ) int {
	// write code here
	if x<=1{
		return x
	}
	l:=1
	r:=x
	for l<r{
		mid := l+((r-l+1)>>1) //这样是为了向上取整
		if mid*mid<=x{
			l = mid
		}else {
			r = mid-1
		}
	}
	return l
}