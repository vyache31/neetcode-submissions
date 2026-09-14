func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var pNext1, pNext2 *ListNode
	p1, p2 := list1, list2
	var prev, head *ListNode
	if p1 == nil && p2 == nil {
		return prev
	} else if p1 == nil && p2 != nil {
		return p2
	} else if p2 == nil && p1 != nil {
		return p1
	}
	if p1.Val > p2.Val {
		prev = p2
		head = p2
	} else {
		prev = p1
		head = p1
	}
	for p1 != nil && p2 != nil {
		pNext1 = p1.Next
		pNext2 = p2.Next
		if p1.Val > p2.Val {
			prev.Next = p2
			prev = p2
			p2 = pNext2
		} else if p2.Val > p1.Val {
			prev.Next = p1
			prev = p1
			p1 = pNext1
		} else {
			prev.Next = p1
			p1.Next = p2
			prev = p2
			p1 = pNext1
			p2 = pNext2
		}
	}
	for p1 == nil && p2 != nil {
		pNext2 = p2.Next
		prev.Next = p2
		prev = p2
		p2 = pNext2
	}
	for p2 == nil && p1 != nil {
		pNext1 = p1.Next
		prev.Next = p1
		prev = p1
		p1 = pNext1
	}
	return head
}