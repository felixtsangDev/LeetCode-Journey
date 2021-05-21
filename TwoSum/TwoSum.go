package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	temp := ListNode{Val: 0, Next: nil}
	result.Val = 0
	for l1.Next != nil {
		if l2.Next != nil {
			temp.Val = l1.Val + l2.Val
			if temp.Val >= 10 {
				temp.Val -= 10
				result.Val = temp.Val
				if l1.Next != nil {
					l1.Next.Val += 1
				} else {
					l2.Next.Val += 1
				}
				newNode := ListNode{Val: 0, Next: nil}
				temp.Next = &newNode
			}
			l2 = l2.Next
		}
		l1 = l1.Next
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
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	l1.print()
	l2.print()
	result := addTwoNumbers(&l1, &l2)
	fmt.Println(result)
}
