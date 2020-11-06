package main

func main() {
	
}
/*
重排列表
1，2，3，4....,n-1,n
1,n,2,n-1...
 */
func reorderList( head *ListNode )  {
	// write code here
	if head==nil||head.Next==nil{
		return
	}
	//通过快慢指针找到这个链表的中点
	slow,fast:=head,head
	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	mid:=slow.Next
	slow.Next=nil
	//反转链表（后半段）
	var pre *ListNode=nil
	for mid!=nil{
		tmp:=mid.Next
		mid.Next=pre
		pre=mid
		mid=tmp
	}
	//合并链表（前半段+反序后的后半段）
	slow=head
	for slow!=nil&&pre!=nil{
		tmp1:=slow.Next
		tmp2:=pre.Next
		slow.Next=pre
		pre.Next=tmp1
		slow=tmp1
		pre=tmp2
	}
}