package main

func main() {
	
}
func maxProfit( prices []int ) int {
	// write code here
	if len(prices) == 0{
		return 0
	}
	res := 0
	min := prices[0]
	for _,val := range prices[1:]{
		if min > val{
			min = val
		}
		if res < val-min{
			res = val-min
		}
	}
	return res
}
