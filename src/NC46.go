package main

import (
	"sort"
)

/*
经典backtry题
加起来和为目标值的组合
 */
func main() {
	combinationSum2([]int{1,1},1)
}

func combinationSum2( num []int ,  target int ) [][]int {
	// write code here
	var path []int
	var res [][]int

	sort.Ints(num)

	dfs(num,0,target,&path,&res)

	return res
}


func dfs(num []int, begin,target int, path *[]int, res *[][]int) {
	if target == 0 {
		tmp:= make([]int,len(*path))
		copy(tmp,*path)
		*res =append(*res, tmp)
		return
	}

	for i:= begin; i < len(num); i++ {
		//剪枝
		if target - num[i] < 0 {
			break
		}
		if i > begin && num[i-1] == num[i] {
			continue
		}
		*path = append(*path, num[i])
		dfs(num,i+1,target-num[i],path,res)
		*path = (*path)[:len(*path)-1]
	}
}