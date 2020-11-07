package main

func main() {
	
}
func maxProfit134( prices []int ) int {
	// write code here
	if len(prices)==0{
		return 0
	}
	res := 0
	for i:=1;i<len(prices);i++{
		pro := prices[i]-prices[i-1]
		if pro>0{
			res += pro
		}
	}
	return res
}