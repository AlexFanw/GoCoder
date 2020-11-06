package main

func main() {
	
}

func ReverseList( pHead *ListNode ) *ListNode {
	// write code here
	if pHead == nil{
		return nil
	}
	var pre *ListNode
	cur := pHead
	for cur!=nil{
		tmp := cur.Next // 保存下一个指针
		cur.Next = pre // 把当前指针指向前一个
		pre = cur // pre移动到cur
		cur = tmp // cur移动到nxt
	}
	return pre
}
