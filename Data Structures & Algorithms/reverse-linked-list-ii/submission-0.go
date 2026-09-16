/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	fakeHead := ListNode{Next: head}
	leftPrev, current := &fakeHead, head
	for i := 0; i < left-1; i++ {
		leftPrev, current = current, current.Next
	}
	var prev *ListNode
	for i := 0; i < right-left+1; i++ {
		tempCurNext := current.Next
		current.Next, prev = prev, current
		current = tempCurNext
	}
	leftPrev.Next.Next = current
	leftPrev.Next = prev
	return fakeHead.Next
}
