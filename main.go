package main

import (
	"fmt"
	list "struct-data/list"
)

func main() {
	//functions.PerfectNumber(6)
	//functions.SubtNumber(1656, 1900)
	//functions.NumberGuay(10)
	//functions.OpenAndClose("{{{}}}", "{")
	//functions.Palindrome("anilina")
	//functions.Palindrome2("anilina")
	//functions.ArrayMethods()
	//functions.Hanoi(4, 1, 3, 2)
	//functions.Recursive()
	//functions.MatrixPow()
	//functions.ThreeInLine()

	l := list.New()
	fmt.Println("######################### Add #############################")
	l.Add(l.Item("Agua", 4))
	l.Add(l.Item("Chocolate", 3))
	l.Add(l.Item("Pizza", 1))
	l.Add(l.Item("Cocacola", 3))
	l.Push(l.Item("Test", 5), 2)
	l.All()
	fmt.Println("######################### Find #############################")
	fmt.Println("item: ", l.Find(3))
	fmt.Println("item first: ", l.First())
	fmt.Println("item last: ", l.Last())
	fmt.Println("######################### Delete #############################")
	l.Delete(2)
	l.All()
}
