package main

func main() {
	
}
/*
判断给定的链表中是否有环? 快慢指针
 */
func hasCycle( head *ListNode ) bool {
	// write code here
	if head == nil || head.Next == nil{
		return false
	}
	fast := head.Next
	slow := head
	for ;;{
		if fast==slow{
			return true
		}
		if fast.Next==nil || fast.Next.Next ==nil{
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
}