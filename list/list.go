package list

import "fmt"

type List struct{}

type Item struct {
	name     string
	quantity int
	next     *Item
}

var root *Item = nil

func New() *List {
	return &List{}
}

func (l *List) All() {
	var item *Item = root

	for i := 1; item != nil; i++ {
		fmt.Printf("%d- Article: %s, quantity: %d \n", i, item.name, item.quantity)
		item = item.next
	}
}

func (l *List) Insert(name string, quantity int) {
	var item Item
	item.name = name
	item.quantity = quantity
	item.next = nil

	if root == nil {
		root = &item
		return
	}

	var pivot *Item
	pivot = root

	for {
		if pivot.next == nil {
			pivot.next = &item
			break
		}

		pivot = pivot.next
	}
}

func (l *List) Delete(index int) {
	if index < 0 || index > l.Size()-1 {
		return
	}

	if index == 0 {
		root = root.next
		return
	}

	var position int = 1
	var item *Item = root

	for {
		if position == index {
			if item.next.next == nil {
				item.next = nil
			} else {
				item.next = item.next.next
			}

			break
		}

		if item != nil {
			item = item.next
		}

		position++
	}
}

func (l *List) Size() int {
	var item *Item = root
	index := 1

	if item == nil {
		return 0
	}

	for item.next != nil {
		index++
		item = item.next
	}

	return index
}
