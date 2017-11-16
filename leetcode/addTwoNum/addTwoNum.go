package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * getListNodeVal: get List Node Value
 * node: current node of linkedlist to retrieve
 * return: next node in the list and the val of the node, or nil and 0
 */
func getListNodeVal(node *ListNode) (*ListNode, int) {
	if node == nil {
		return nil, 0
	}
	return node.Next, node.Val
}

/**
 * createSumNode: create and appends a node that stores just the onesPlace to a linkedlist
 * node: node of linkedlist or nil
 * num: num to split into onesPlace and remainder
 * return: last node of linkedlist and remainder
 */
func createSumNode(node *ListNode, num int) (*ListNode, int) {
	var onesPlace = num % 10
	var remainder = num / 10
	var tmpNode = &ListNode{onesPlace, nil}
	if node != nil {
		node.Next = tmpNode
	}
	return tmpNode, remainder
}

/**
 * addTwoNumbers: sum two numbers represented by a singly linked list. Digits are reversed
 * l1: pointer to a reversed number (each node holds one digit)
 * l2: pointer to a reversed number (each node holds one digit)
 * return: pointer to the sum of the two numbers
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumTwoNumList *ListNode
	var l1Num, l2Num, remainder int
	l1, l1Num = getListNodeVal(l1)
	l2, l2Num = getListNodeVal(l2)
	sumTwoNumList, remainder = createSumNode(nil, l1Num+l2Num)
	var node = sumTwoNumList
	for l1 != nil || l2 != nil {
		l1, l1Num = getListNodeVal(l1)
		l2, l2Num = getListNodeVal(l2)
		node, remainder = createSumNode(node, l1Num+l2Num+remainder)
	}
	if remainder != 0 {
		createSumNode(node, remainder)
	}
	return sumTwoNumList
}

func appendNode(node *ListNode, num int) (*ListNode, int) {
	var onesPlace = num % 10
	var restOfNum = num / 10
	var tmpNode = &ListNode{onesPlace, nil}
	if node != nil {
		node.Next = tmpNode
	}
	return tmpNode, restOfNum
}

func intToLinkedList(num int) *ListNode {
	//creating head node
	var head, node *ListNode
	head, num = appendNode(nil, num)
	node = head
	for num != 0 {
		//appending to the list
		node, num = appendNode(node, num)
	}
	return head
}

func printLinkedList(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
	fmt.Println("\n\n\n ")
}
func main() {
	var num1 = 123
	var num2 = 78
	var l1 = intToLinkedList(num1)
	var l2 = intToLinkedList(num2)
	printLinkedList(l1)
	printLinkedList(l2)
	printLinkedList(addTwoNumbers(l1, l2))
}
