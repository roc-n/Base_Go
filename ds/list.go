package ds

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 定义节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 定义链表
type LinkedList struct {
	Head  *ListNode
	End   *ListNode
	Items int
}

// 创建链表
func CreateList() *LinkedList {
	return &LinkedList{
		Head:  nil,
		End:   nil,
		Items: 0,
	}
}

// 添加节点
func (l *LinkedList) AddNode() {
	newNode := &ListNode{Next: nil}
	if l.Items == 0 {
		l.Head = newNode
		l.End = newNode
	} else {
		l.End.Next = newNode
		l.End = newNode
	}
	l.Items++
}

// 添加数据
func (l *LinkedList) Add() {
	l.AddNode()
	l.GetInfo()
}

// 获取节点信息
func (l *LinkedList) GetInfo() {
	fmt.Println("Enter the data:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data, _ := strconv.Atoi(scanner.Text())
	l.End.Val = data
}

// 判断链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.Items == 0
}

// 删除节点
func (l *LinkedList) Delete() {
	fmt.Println("Enter the data to delete:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data, _ := strconv.Atoi(scanner.Text())

	currentNode := l.Head
	var preNode *ListNode

	for currentNode != nil {
		if currentNode.Val == data {
			break
		}
		preNode = currentNode
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		fmt.Println("Not found, can't delete.")
		return
	}

	if currentNode == l.Head {
		l.Head = currentNode.Next
	} else if currentNode == l.End {
		preNode.Next = nil
		l.End = preNode
	} else {
		preNode.Next = currentNode.Next
	}
	l.Items--
	fmt.Println("Done")
}

// 查找数据
func (l *LinkedList) Search() {
	fmt.Println("Enter the data to search:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data, _ := strconv.Atoi(scanner.Text())

	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Val == data {
			fmt.Println("Found it.")
			return
		}
		currentNode = currentNode.Next
	}
	fmt.Println("Not found.")
}

// 清空链表
func (l *LinkedList) MakeEmpty() {
	l.Head = nil
	l.End = nil
	l.Items = 0
}

// 销毁链表
func (l *LinkedList) Dispose() {
	l.MakeEmpty()
}

func ListDemo() {
	l := CreateList()
	for range 10 {
		l.Add()
	}
	l.Delete()
	l.Search()
	l.Dispose()
}
