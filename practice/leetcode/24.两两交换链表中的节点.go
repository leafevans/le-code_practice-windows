/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next.Next

		prev.Next = a.Next
		prev.Next.Next = a
		a.Next = b

		prev = a
	}

	return dummy.Next
}

// @lc code=end

