package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalfStart := reverseList(slow)
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	slow.Next = reverseList(secondHalfStart)

	return result

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next // 暂存下一个节点
		curr.Next = prev      // 改变指针方向
		prev = curr           // prev移动到当前节点
		curr = nextTemp       // curr移动到下一个节点
	}
	return prev
}

func main() {
	// var s []int
	// s = append(s, 1)
	// s = append(s, 2)
	// s = append(s, 3)
	// s = append(s, 3)
	// s = append(s, 2)
	// s = append(s, 1)
	// fmt.Printf("1:%d\n", s[0])
	// fmt.Printf("6:%d\n", s[len(s)-1])
	test1 := &ListNode{Val: 10}
	test1.Next = &ListNode{Val: 11}
	test1.Next.Next = &ListNode{Val: 11}
	test1.Next.Next.Next = &ListNode{Val: 10}
	test1.Next.Next.Next.Next = nil
	// for test1 != nil {
	// 	fmt.Println(test1.Val)
	// 	test1 = test1.Next
	// }
	_ = isPalindrome(test1)
	for test1 != nil {
		fmt.Println(test1.Val)
		test1 = test1.Next
	}
}
