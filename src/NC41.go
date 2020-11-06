package main

import "fmt"

func main() {
	arr := []int{2,3,4,5}
	fmt.Println(maxLength(arr))
	s := "aabaab!bb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func maxLength( arr []int ) int {
	// write code here
	maxset := make(map[int] int)
	res := 0
	slow:=0
	for fast:=0; fast<len(arr);fast++{
		if idx,ok := maxset[arr[fast]];ok{//存在与数组中
			if idx>=slow{
				slow = maxset[arr[fast]]+1 //slow移动到重复数的后一位
			}
		}
		if fast-slow+1 > res{
			res = fast-slow+1 //判断现在的最新状态
		}
		maxset[arr[fast]] = fast //重复数的位置更新为最新状态
	}
	return res
}

func lengthOfLongestSubstring(arr string) int {
	// write code here
	maxset := make(map[byte] int)
	res := 0
	slow:=0
	for fast:=0; fast<len(arr);fast++{
		if idx,ok := maxset[arr[fast]];ok{//如果fast存在与数组中
			if idx>=slow{
				slow = idx+1//slow移动到与fast重复的数的后一位
			}
		}
		maxset[arr[fast]] = fast //fast的位置更新为最新状态
		if fast-slow+1 > res{
			res = fast-slow+1 //判断现在的最新状态
		}

	}
	return res
}