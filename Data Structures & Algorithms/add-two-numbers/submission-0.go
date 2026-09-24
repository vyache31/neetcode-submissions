/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	sum := []int{}
	ost := 0
	for i := 0; i < max(len(stack1), len(stack2)); i++ {
		add := 0 + ost
        ost = 0
		if i+1 <= len(stack1) {
			add += stack1[i]
		}
		if i+1 <= len(stack2) {
			add += stack2[i]
		}
		if add > 9 {
			ost = 1
			add -= 10
		}
		sum = append(sum, add)
	}
	if ost != 0 {
		sum = append(sum, ost)
	}
	prev := &ListNode{}
	headPtr := prev
	for i := 0; i < len(sum); i++ {
		curr := &ListNode{
			Val:  sum[i],
			Next: nil,
		}
		prev.Next = curr
		prev = curr
	}
	return headPtr.Next
}