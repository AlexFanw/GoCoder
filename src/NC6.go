package main

import "math"

/*
二叉树的最大路径和 ，读题比写题难
 */
func main() {
	
}
var res int
func maxPathSum( root *TreeNode ) int {
	// write code here
	res = math.MinInt64
	getMax(root)
	return res
}
func getMax(root * TreeNode) int{
	if root==nil{
		return 0
	}
	leftmax := listMax([]int{getMax(root.Left),0}) // 保证了大于等于0
	rightmax := listMax([]int{getMax(root.Right),0}) // 保证了大于等于0
	nowmax := root.Val + leftmax + rightmax// 以当前结点为根的最大和
	if nowmax>res{
		res = nowmax
	}
	return listMax([]int{leftmax,rightmax})+root.Val //但是如果要递归给上层用的话，只能选一条路径噢！！！！这点非常重要
}

func listMax(l []int) int{//返回最大值，如果都是负数则返回0
	mx := 0
	for _,v := range l{
		if mx < v{
			mx = v
		}
	}
	return mx
}