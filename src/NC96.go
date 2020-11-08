package main

func main() {
	
}
func isPail( head *ListNode ) bool {
	// write code here
	dummy := &ListNode{Next:head}
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	// 反转后半部分链表
	var pre, next *ListNode
	cur := mid
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	tail := pre
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}