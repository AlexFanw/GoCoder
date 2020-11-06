package main

func main() {
	
}
/*
二叉树根节点到叶子节点的所有路径和
输入{1,0}
返回 10
 */
func sumNumbers( root *TreeNode ) int {
	prt := 0
	return getSum(root, prt)
}

func getSum( root *TreeNode,prt int)int{
	if root == nil{
		return 0
	}
	prt = prt*10+root.Val
	leftline := 0
	rightline := 0
	if root.Left!=nil {
		leftline = getSum(root.Left, prt) // 左子树的总收益
	}
	if root.Right!=nil{
		rightline = getSum(root.Right, prt) // 右子树的总收益
	}
	if root.Left==nil && root.Right==nil{
		return prt //没有左右子树则返回当前收益
	}
	return leftline+rightline
}