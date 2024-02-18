package main

import (
	"com.github/dukendev/basics/linkedlist"
)

func main() {
	linkedListDemo()
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

	secondList.AppendAtStart(&firstList)
	secondList.Show()
	firstList.AppendAtStart(&linkedlist.LinkedList[int]{})
	firstList.Show()
}
