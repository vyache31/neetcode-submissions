/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = prev
		prev = curNode
		curNode = nextNode
	}
	return prev
}
