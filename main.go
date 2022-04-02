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
	l.Insert("Agua", 4)
	l.Insert("Chocolate", 3)
	l.Insert("Pizza", 1)
	l.Insert("Cocacola", 1)
	l.All()
	fmt.Println("######################################################")
	l.Delete(2)
	l.All()
}
