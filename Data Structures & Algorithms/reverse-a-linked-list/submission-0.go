/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	curNode := head
	for {
		if curNode == nil {
			break
		}
		nextNode := curNode.Next
		curNode.Next = prev
		prev = curNode
		curNode = nextNode
	}
	return prev
}
