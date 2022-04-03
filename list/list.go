package list

import "fmt"

type List struct{}

type Node struct {
	Name     string
	Quantity int
	Next     *Node
}

var root *Node = nil

func New() *List {
	return &List{}
}

func (l *List) All() {
	var item *Node = root

	for i := 1; item != nil; i++ {
		fmt.Printf("%d- Article: %s, quantity: %d \n", i, item.Name, item.Quantity)
		item = item.Next
	}
}

func (l *List) Size() int {
	var item *Node = root
	index := 1

	if item == nil {
		return 0
	}

	for item.Next != nil {
		index++
		item = item.Next
	}

	return index
}

func (l *List) Item(name string, quantity int) Node {
	return Node{
		Name:     name,
		Quantity: quantity,
		Next:     nil,
	}
}

func (l *List) Add(item Node) {
	if root == nil {
		root = &item
		return
	}

	var pivot *Node
	pivot = root

	for {
		if pivot.Next == nil {
			pivot.Next = &item
			break
		}

		pivot = pivot.Next
	}
}

func (l *List) Push(item Node, index int) {
	var position int = 1
	var pivot *Node

	if index < position || index > l.Size() {
		return
	}

	if index == position {
		item.Next = root
		root = &item
		return
	}

	pivot = root

	for {
		// -1 hace que se salga una en el item anterior
		if position == index-1 {
			break
		}

		pivot = pivot.Next
		position++
	}

	item.Next = pivot.Next
	pivot.Next = &item
}

func (l *List) Find(index int) Node {
	item := root
	position := 1

	if index < position || index > l.Size() {
		return Node{}
	}

	for {
		if index == position {
			break
		}

		item = item.Next
		position++
	}

	return *item
}

func (l *List) First() Node {
	if l.Size() == 0 {
		return Node{}
	}

	return *root
}

func (l *List) Last() Node {
	item := root
	position := 1

	if l.Size() == position {
		return *item
	}

	for position < l.Size() {
		item = item.Next
		position++
	}

	return *item
}

func (l *List) Delete(index int) {
	positionRoot := 1

	if index < positionRoot || index > l.Size() {
		return
	}

	if index == positionRoot {
		root = root.Next
		return
	}

	destroyItem(index, positionRoot+1, root)
}

func destroyItem(index int, position int, item *Node) {
	if position == index {
		if item.Next.Next == nil {
			item.Next = nil
		} else {
			item.Next = item.Next.Next
		}

		return
	}

	if item != nil {
		item = item.Next
	}

	destroyItem(index, position+1, item)
}
