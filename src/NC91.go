package main

import "fmt"

func main() {
	inp := []int{2,1,5,3,6,4,8,9,7}
	LIS(inp)
}
/*
能想到动态规划就很不错了。
dp[i]表示以第i个元素结尾的最长子序列
dp[i] = max(dp[i],dp[j]+1),这个j表示所有比第i个元素小的元素
 */

/*
有一种比较巧的解法可以看看：用二分法，更快
https://zhuanlan.zhihu.com/p/78224512
 */
func LIS( arr []int ) []int {
	// write code here
	if len(arr) == 0{
		return nil
	}
	dp := []int{}
	//res := 0
	for arrindex, _:= range arr{
		max := 1
		for index, val := range arr[0:arrindex]{
			if val < arr[arrindex]{
				fmt.Println(arr[arrindex])
				if(max < dp[index]+1){
					max = dp[index]+1
				}
			}
		}
		dp = append(dp, max)
	}
	fmt.Println(dp)
	return dp;
}