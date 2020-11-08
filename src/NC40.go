package main

func main() {
	
}

func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
	// write code here
	head1, head2 = reverse(head1), reverse(head2)
	tmp := 0
	prev := &ListNode{}
	current := prev
	for head1 != nil  || head2 != nil || tmp > 0 {
		if head1 != nil {
			tmp += head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			tmp += head2.Val
			head2 = head2.Next
		}
		current.Next = &ListNode{Val: tmp % 10}
		current = current.Next
		tmp /= 10
	}
	return reverse(prev.Next)
}

func reverse40(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}
