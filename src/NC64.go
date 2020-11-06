package main

func main() {
	
}
/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
 */

/*
先中序遍历，然后把结果存在一个list里，然后再重定向
 */
func Convert( pRootOfTree *TreeNode ) *TreeNode {
	// write code here
	if pRootOfTree==nil{
		return nil
	}
	nodelist := []*TreeNode{}
	nodelist = treeToList(pRootOfTree, nodelist)
	if len(nodelist)==1{
		return pRootOfTree
	}
	for k,v := range nodelist{
		if k>0 && k<len(nodelist)-1{
			v.Left = nodelist[k-1]
			v.Right = nodelist[k+1]
		}
		if k==0{
			v.Left = nil
			v.Right = nodelist[1]
		}
		if k==len(nodelist)-1{
			v.Left = nodelist[k-1]
			v.Right = nil
		}
	}
	return nodelist[0]
}

func treeToList( root *TreeNode, nodelist []*TreeNode)[]*TreeNode{
	if root==nil{
		return nodelist
	}
	nodelist = treeToList(root.Left, nodelist)
	nodelist = append(nodelist, root)
	nodelist = treeToList(root.Right, nodelist)
	return nodelist
}