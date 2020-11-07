package main

func main() {
	
}
func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	var m = make(map[*ListNode]bool, 0)
	var p1, p2 *ListNode

	for p1 = pHead1; p1 != nil; p1 = p1.Next {
		m[p1] = true
	}
	for p2 = pHead2; !m[p2] && p2 != nil; p2 = p2.Next {}
	return p2
}
