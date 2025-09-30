package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Wahid")
	data.PushBack("Rizka")
	data.PushBack("Fathurrohman")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // wahid

	next := head.Next() // Rizka
	fmt.Println(next.Value)

	next = next.Next() // Fathurrohman
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
