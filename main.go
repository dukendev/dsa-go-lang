package main

import (
	"fmt"

	"com.github/dukendev/basics/algorithms"
	"com.github/dukendev/basics/linkedlist"
)

func main() {
	binarySearchDemo()
}

func binarySearchDemo() {
	list := []int{1, 2, 3, 4, 5}
	r := algorithms.BinarySearch(list, 6)
	fmt.Print(r)
}

func linearSearchDemo() {
	list := []int{1, 2, 3, 4, 5}
	r := algorithms.LinearSearch(list, 3)
	fmt.Print(r)
}

func linkedListDemo() {
	q := linkedlist.CreateNode(1)
	n := linkedlist.CreateNode(2)
	m := linkedlist.CreateNode(3)
	o := linkedlist.CreateNode(4)
	p := linkedlist.CreateNode(5)
	firstList := linkedlist.LinkedList[int]{}
	firstList.Add(q)
	firstList.Add(n)
	firstList.Add(m)
	firstList.Add(o)
	firstList.Add(p)
	firstList.Show()

	secondList := linkedlist.LinkedList[int]{}
	secondList.Add(linkedlist.CreateNode[int](10))
	secondList.Add(linkedlist.CreateNode[int](20))

	secondList.AppendAtEnd(&firstList)
	secondList.Show()
}
