package main

import "fmt"

func main() {
	in := [][]int{{1,   3,  5,  9}, {10, 11, 12, 30}, {230, 300, 350, 500}}
	fmt.Println(spiralOrder(in))
}

func spiralOrder( matrix [][]int ) []int {
	// write code here
	res := []int{}
	m := len(matrix) //3
	n := len(matrix[0]) //4
	sum := m*n
	status := 0 // 0 向右 1 向下 2 向左 3 向上
	x := 0
	y := 0
	step := 1
	for i:=0;i<sum;i++{
		res = append(res, matrix[y][x])
		//判断是否需要调头
		switch status {
		case 0,2:
			if step==n {
				step = 0
				m-=1
				status=(status+1)%4
			}
		case 1,3:
			if step==m {
				step = 0
				n-=1
				status=(status+1)%4
			}

		}
		// 判断需要往哪个方向走一步
		switch status {
		case 0:
			x+=1
		case 1:
			y+=1
		case 2:
			x-=1
		case 3:
			y-=1
		}
		step+=1
	}
	return res
}