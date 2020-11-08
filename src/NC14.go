package main

func main() {
	
}

func zigzagLevelOrder( root *TreeNode ) [][]int {
	// write code here
	if root == nil{
		return [][]int{}
	}
	ret := [][]int{}
	stack := []*TreeNode{root}
	level := 1
	for {
		tmp := []int{}
		next := []*TreeNode{}
		for _, node :=range stack{
			if node != nil{
				next = append(next, node.Left, node.Right)
				if level > 0{
					tmp = append(tmp, node.Val)
				}else{
					tmp = append([]int{node.Val}, tmp...)
				}
			}
		}
		if len(next) >0{
			stack = next
		}else{
			break;
		}
		ret = append(ret, tmp)
		if level > 0{
			level = -1
		}else{
			level = 1
		}
	}
	return ret
}