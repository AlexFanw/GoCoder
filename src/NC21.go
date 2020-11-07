package main

func main() {
	
}

func reverseBetween( head *ListNode ,  m int ,  n int ) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	// write code here
	res:=new(ListNode)
	res.Next=head
	pre,end:=res,res
	for i:=0;i<m-1;i++{
		pre=pre.Next
	}
	start:=pre.Next
	for i:=0;i<n;i++{
		end=end.Next
	}
	next:=end.Next
	end.Next=nil

	pre.Next = reverse(start)
	start.Next = next

	return res.Next
}

func reverse(head *ListNode)*ListNode{
	var pre *ListNode
	for head!=nil{
		next:=head.Next
		head.Next=pre
		pre=head
		head=next
	}
	return pre
}
