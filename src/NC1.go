package main

import "fmt"

func main() {
	fmt.Println(solve1("1","789"))
}
func solve1( s string ,  t string ) string {
	// write code here
	ret := []byte{}
	var carryBit byte
	for i,j:=len(s)-1,len(t)-1;i >=0 || j>=0 ||carryBit >0;i,j=i-1,j-1{
		var x,y byte
		if i>=0{
			x = s[i] - '0'
		}
		if j>=0{
			y = t[j] -'0'
		}
		cur := x+y+carryBit
		carryBit = cur/10
		fmt.Println(cur)
		ret = append([]byte{cur%10+'0'}, ret...)
	}
	return string(ret)
}