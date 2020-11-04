package main

import "fmt"

func main() {
	inp := []int{1,2,8,6,4}
	fmt.Println(LIS(inp))
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
				if(max < dp[index]+1){
					max = dp[index]+1
				}
			}
		}
		dp = append(dp, max)
	}
	//这上面可以得出以每一个元素结尾的最大递增长度

	//找出最大的子序列，而且一定是最靠后的那个才是我们要找的
	max := 0
	maxindex := 0
	for index,val := range dp{
		if val>=max{
			max = val
			maxindex = index
		}
	}

	//恢复出子序列
	res := []int{arr[maxindex]}
	for i:=maxindex-1;i>=0;i--{
		if arr[i] < arr[maxindex] && dp[i] == max-1{
			res = append([]int{arr[i]},res...)
			max-=1
		}
	}
	return res;
}