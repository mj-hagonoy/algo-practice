package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root *ListNode
	var currNode *ListNode

	for list1 != nil && list2 != nil {
		val := 0
		if list1.Val <= list2.Val {
			val = list1.Val
			list1 = list1.Next

		} else {
			val = list2.Val
			list2 = list2.Next
		}

		root, currNode = _append(root, currNode, &ListNode{Val: val, Next: nil})
	}

	for list1 != nil {
		root, currNode = _append(root, currNode, &ListNode{Val: list1.Val, Next: nil})
		list1 = list1.Next

	}

	for list2 != nil {
		root, currNode = _append(root, currNode, &ListNode{Val: list2.Val, Next: nil})
		list2 = list2.Next
	}

	return root
}

func _append(root *ListNode, currNode *ListNode, node *ListNode) (*ListNode, *ListNode) {
	if root == nil {
		root = node
	} else {
		currNode.Next = node
	}
	currNode = node
	return root, currNode
}

/*
MERGE TWO SORTED LIST
===
Result: ACCEPTED
Runtime: 2 ms, faster than 36.30% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.7 MB, less than 28.07% of Go online submissions for Merge Two Sorted Lists.
====
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

*/
