package main
/*
前序中序还原一颗二叉树
 */
func main() {
	pre := []int{3,9,20,15,17}
	ino := []int{9,3,15,20,7}
	buildTree(pre,ino)

}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0{
		return nil
	}
	root := TreeNode{}
	root.Val = preorder[0]
	index := 0
	for ;;index++{
		if inorder[index]==preorder[0]{
			break
		}
	}//分割inorder
	root.Left = buildTree(preorder[1:1+index],inorder[:index])
	root.Right = buildTree(preorder[1+index:],inorder[index+1:])
	return &root
}