package main

import (
	"fmt"
)

func main() {
	in := [][]int{{1,   3,  5,  9}, {10, 11, 12, 30}, {230, 300, 350, 500}}
	searchMatrix(in, 9)
}
/*
请写出一个高效的在m*n矩阵中判断目标值是否存在的算法，矩阵具有如下特征：
每一行的数字都从左到右排序
每一行的第一个数字都比上一行最后一个数字大
例如：
对于下面的矩阵
 */

/*
输入：
[
    [1,   3,  5,  9],
    [10, 11, 12, 30],
    [230, 300, 350, 500]
] 3
返回：true
 */

func searchMatrix( matrix [][]int ,  target int ) bool {
	// write code here
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1]{
		return false
	}
	l:=0
	r:=len(matrix)-1
	for r-l>1{
		mid := (r + l ) / 2

		fmt.Println("mid",mid)
		if matrix[mid][0] > target{
			r = mid
			fmt.Println("r",mid)
		}else if matrix[mid][len(matrix[mid])-1] < target {
			l = mid+1
		}else if matrix[mid][0] < target {
			l = mid
			//fmt.Println(mid)
		} else if matrix[mid][0] == target{
			//fmt.Println(target)
			return true
		}else {
			break
		}
	}
	v := matrix[l]
	r = len(matrix[l])-1
	//fmt.Println(v)
	l = 0
	for l<=r{
		mid := (r+l)/2
		if v[mid] > target{
			r = mid-1
		}else if v[mid] < target{
			l = mid+1
		} else {
			//fmt.Println(v[mid])
			return true
		}
	}
	return false
}