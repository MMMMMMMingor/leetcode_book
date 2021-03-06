/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	quick := head
	slow := head
	hasCycle := false

	for quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

		if quick == slow {
			hasCycle = true
			break
		}
	}

	if hasCycle == false{
		return nil
	}

	find := head

	for find != slow{
		find = find.Next
		slow = slow.Next
	}

	return slow
}
// @lc code=end

