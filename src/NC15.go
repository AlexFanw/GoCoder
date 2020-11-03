package main
import . "nc_tools"
func main() {
	
}

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 * 给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
 */

/*
使用双队列来解决
 */
func levelOrder( root *TreeNode ) [][]int {
	// write code here
	if root == nil{
		return nil
	}
	var result [][]int
	nodes := []*TreeNode{root}
	for len(nodes)>0{
		level := make([]int,0)
		checknodes := nodes[:]
		nodes = make([]*TreeNode,0)
		for _, node := range checknodes{
			level = append(level, node.Val)
			if node.Left!=nil{
				nodes = append(nodes, node.Left)
			}
			if node.Right!=nil{
				nodes = append(nodes, node.Right)
			}
		}
		result = append(result, level)
	}
	return result;
}