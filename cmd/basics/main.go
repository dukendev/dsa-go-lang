package main

import (
	"com.github/dukendev/basics/linkedlist"
)

func main() {
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
	firstList.Remove(3)
	firstList.Show()
	firstList.Search(3)
	firstList.AddAtStart(0)
	firstList.Show()
	firstList.RemoveAtStart()
	firstList.Show()
}
