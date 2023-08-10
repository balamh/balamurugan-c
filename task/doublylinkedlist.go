package main

import (
	"fmt"
)

type node struct {
	value    int
	previous *node
	next     *node
}
type doublelinkedlist struct {
	head   *node
	tail   *node
	length int
}

func (d *doublelinkedlist) append(val int) {
	newlist := &node{val, nil, nil}
	if d.length == 0 {
		d.head = newlist
		d.tail = newlist
		d.length++
	} else {
		d.tail.previous = d.tail
		d.tail.next = newlist
		d.tail = newlist
	}
}
func (d *doublelinkedlist) printfrwd() {
	temp := d.head
	if d.length == 0 {
		fmt.Println("No element Present")
	} else {
		fmt.Println("The Elements in List Forward")
		for temp != nil {
			fmt.Println(temp.value)
			temp = temp.next
		}
	}
}
func (d *doublelinkedlist) reverse() {
	temp1 := d.tail
	if d.length == 0 {
		fmt.Println("No element Present")
	} else {
		fmt.Println("The Elements in List Backward")
		for temp1 != nil {
			fmt.Println(temp1.value)
			temp1 = temp1.previous
		}
	}
}
func main() {
	dlist := doublelinkedlist{}
	dlist.append(5)
	dlist.append(6)
	dlist.append(3)
	dlist.printfrwd()
	dlist.reverse()
}
