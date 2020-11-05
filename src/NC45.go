package main
/*
分别按照二叉树先序，中序和后序打印所有的节点。
 */
func main() {
	
}
/*
递归写法，很快。一目了然
 */
func threeOrders( root *TreeNode ) [][]int {
	// write code here
	pre := []int{}
	cur := []int{}
	bak := []int{}
	pre = preOrders(root, pre)
	cur = curOrders(root, cur)
	bak = bakOrders(root, bak)
	return [][]int{pre,cur,bak}
}

func preOrders( root *TreeNode, pre []int) []int{
	if root==nil{
		return pre
	}
	pre = append(pre, root.Val)
	pre = preOrders(root.Left, pre)
	pre = preOrders(root.Right, pre)
	return pre
}
func curOrders( root *TreeNode, cur []int) []int{
	if root==nil{
		return cur
	}
	cur = curOrders(root.Left, cur)
	cur = append(cur, root.Val)
	cur = curOrders(root.Right, cur)
	return cur
}
func bakOrders( root *TreeNode, bak []int) []int{
	if root==nil{
		return bak
	}
	bak = bakOrders(root.Left, bak)
	bak = bakOrders(root.Right, bak)
	bak = append(bak, root.Val)
	return bak
}