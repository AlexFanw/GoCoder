package main

func main() {
	
}

func lowestCommonAncestor( root *TreeNode ,  o1 int ,  o2 int ) int {
	// write code here
	if root==nil{
		return -1
	}
	if root.Val==o1 ||root.Val==o2 {
		return root.Val
	}//已经探测到了目标结点
	right:=lowestCommonAncestor(root.Right ,o1,o2) // 探测右边
	left:=lowestCommonAncestor(root.Left, o1, o2) // 探测左边
	if left!=-1 && right!=-1{ //同时满足左右都不为空，才算真正意义上的父结点
		return root.Val
	}
	if right!=-1{
		return right
	}else if left!=-1{
		return left
	}
	return -1
}

