package main

func main() {
	
}
type TreeNode2 struct {
	Val int
	Left *TreeNode2
	Right *TreeNode2
}

func solve136( xianxu []int ,  zhongxu []int ) []int {
	// write code here
	root := helper(xianxu, zhongxu, 0, len(xianxu)-1, 0, len(zhongxu)-1)
	return transfer(root)
}

func transfer(root *TreeNode2) []int{
	if root == nil {
		return nil
	}
	q := []*TreeNode2{root}
	ans := []int{}
	for len(q) > 0 {
		curVal := 0
		n := len(q)
		for i := 0; i < n; i ++ {
			curNode := q[i]
			curVal = curNode.Val
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
		q = q[n:]
		ans = append(ans, curVal)
	}
	return ans
}

func helper(xianxu []int, zhongxu []int, x1, x2, z1, z2 int) *TreeNode2 {
	if x1 > x2 {
		return nil
	}
	curHead := &TreeNode2{
		Val : xianxu[x1],
	}
	i := 0
	for i = z1; i < z2; i ++ {
		if zhongxu[i] == xianxu[x1] {
			break
		}
	}
	curHead.Left = helper(xianxu, zhongxu, x1+1, x1+i-z1, z1, i-1)
	curHead.Right = helper(xianxu, zhongxu, x1+i-z1+1, x2, i+1, z2)
	return curHead
}
