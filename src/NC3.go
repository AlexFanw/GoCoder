package main

func main() {
	
}
/*
对于一个给定的链表，返回环的入口节点，如果没有环，返回null
拓展：
你能给出不利用额外空间的解法么？
 */
func detectCycle( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}