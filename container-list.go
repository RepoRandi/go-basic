package main

import "container/list"

func main53() {
	data := list.New()

	data.PushBack(1)
	data.PushBack(2)
	data.PushBack(3)

	for e := data.Front(); e != nil; e = e.Next() {
		println(e.Value)
	}
}
