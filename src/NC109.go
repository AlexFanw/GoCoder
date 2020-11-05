package main

func main() {
	
}
/**
 * 判断岛屿数量
 * @param grid char字符型二维数组
 * @return int整型
 */

/*
并查集写法
 */
func solve109( grid [][]byte ) int {
	// write code here
	row := len(grid)
	col := len(grid[0])
	Tree := make([]int, row*col)
	for i:= 0; i<row*col;i++ {
		Tree[i] = -1
	}
	count := 0
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if grid[i][j] == '1' {
				count++
				if i+1 < row && grid[i+1][j]=='1'&&merge(i*col+j,(i+1)*col+j,Tree) {
					count--
				}
				if j+1 < col && grid[i][j+1] == '1' && merge(i*col+j, i*col+j+1, Tree) {
					count--
				}
			}
		}
	}
	return count
}
func merge(a, b int, Tree []int) bool {
	a = findr(a, Tree)
	b = findr(b, Tree)
	if a!=b {
		Tree[a] = b
		return true
	}
	return false
}
func findr(a int, Tree []int) int {
	tmp := a
	for Tree[a]!=-1 {
		a = Tree[a]
	}
	ret := a
	a = tmp
	// 再做一次从节点a到根节点的遍历
	for Tree[a] != -1 {
		t := Tree[a]
		Tree[a] = ret
		a = t
	}
	return a
}

/*
DFS写法 推荐用这个，并查集真的折磨人...
 */

func solve109DFS( grid [][]byte ) int {
	res := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]!='0'{
				res++
				dfsGrid(grid,i,j)
			}
		}
	}
	return res
}

func dfsGrid(grid [][]byte ,i int,j int){
	if i<0||j<0||i>=len(grid)||j>=len(grid[0]){
		return
	}
	if grid[i][j]=='1'{
		grid[i][j] = '0' //把它藏起来
		dfsGrid(grid, i+1,j)
		dfsGrid(grid, i ,j+1)
		dfsGrid(grid, i-1,j)
		dfsGrid(grid, i ,j-1)
		return
	}
}