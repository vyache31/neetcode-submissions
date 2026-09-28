/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := head
	target := head
	curr := head
	var dist int
	for curr != nil {
		if dist == n {
			prev = target
			target = target.Next
			dist--
		}
		curr = curr.Next
		dist++
	}
    if target == head {
        return head.Next
    }
    
	prev.Next = target.Next
	target.Next = nil
	return head
}
