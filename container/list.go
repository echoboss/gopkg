package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	mid := l.PushBack(2)
	back := l.PushBack(3)
	frond := l.PushFront(1)

	fmt.Println(mid.Next().Value)
	fmt.Println(mid.Prev().Value)

	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)

	frond = l.InsertBefore(0,frond)
	back = l.InsertBefore(4,back)

	l.MoveToBack(frond)
	l.MoveToFront(back)

	//printList(l)

	cpl := list.New()
	cpl.PushBack("abc")
	cpl.PushBackList(l)


	l.Init()
	printList(cpl)


}

func printList(l *list.List)  {
	for  e := l.Front(); e != nil;e = e.Next() {
		fmt.Print(e.Value," ")
	}
}
