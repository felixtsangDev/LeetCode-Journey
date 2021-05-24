package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	result = &ListNode{Val: 0, Next: nil}
	temp := result
	var i int = 0
	for l1 != nil {
		if i != 0 {
			temp.Next = &ListNode{Val: 0, Next: nil}
			temp = temp.Next
		}
		if l2 != nil {
			temp.Val = l1.Val + l2.Val
			if temp.Val >= 10 {
				temp.Val -= 10
				if l1.Next != nil {
					l1.Next.Val += 1
				} else {
					l1.Next = &ListNode{Val: 1}
				}
			}
			l1 = l1.Next
			l2 = l2.Next
		} else {
			temp.Val = l1.Val
			if temp.Val >= 10 {
				temp.Val -= 10
				if l1.Next != nil {
					l1.Next.Val += 1
				} else {
					l1.Next = &ListNode{Val: 1}
				}
			}
			l1 = l1.Next
		}
		i++
	}
	for l2 != nil {
		if i != 0 {
			temp.Next = &ListNode{Val: 0, Next: nil}
			temp = temp.Next
		}
		temp.Val = l2.Val
		if temp.Val >= 10 {
			temp.Val -= 10
			if l2.Next != nil {
				l2.Next.Val += 1
			} else {
				l2.Next = &ListNode{Val: 1}
			}
		}
		l2 = l2.Next
	}
	return
}

func (l *ListNode) print() {
	temp := l
	fmt.Println(temp.Val)
	for temp.Next != nil {
		fmt.Println(temp.Next.Val)
		temp = temp.Next
	}
}

func main() {
	// l1 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}
	// l2 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}

	result := addTwoNumbers(&l1, &l2)
	result.print()
}
