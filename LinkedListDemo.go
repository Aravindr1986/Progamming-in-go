package main

import "fmt"

type LinkedList struct {
	val  int
	next *LinkedList
	prev *LinkedList
}

func createSingleConnectedList(startVal int) *LinkedList { //function to create singley linked list

	node := LinkedList{startVal, nil, nil}
	head := &node
	return head
}
func CreateDoubleConnectedList(startVal int) *LinkedList { //function to create singley linked list

	node := LinkedList{startVal, nil, nil}
	head := &node
	node.prev = head
	return head
}
func DisplayList(head *LinkedList) {

	ptr := head
	for ptr != nil {
		fmt.Print(" ", ptr.val)
		ptr = ptr.next
	}
}
func DeleteElement(head *LinkedList, position int) *LinkedList {
	ptr := head
	prev := head
	if position == 0 { //if element is to be entered into the first position
		ptr = head
		head = head.next
		ptr.next = nil
	} else {
		for i := 0; i < position; i++ {
			if ptr.next == nil { //handling for append cases
				break
			}
			prev = ptr
			ptr = ptr.next //iterate through the graph
		}
		if ptr.next == nil {
			prev.next = nil
		} else {
			temp := ptr.next
			ptr.next = ptr.next.next
			temp.next = nil
		}
	}
	return head
}
func InsertElement(head *LinkedList, element int, position int) *LinkedList {

	ptr := head
	node := LinkedList{element, nil, nil}
	if position == 0 { //if element is to be entered into the first position
		node.next = head
		head = &node
	} else { //any other position includeing end
		for i := 0; i < position; i++ {
			if ptr.next == nil { //handling for append cases
				break
			}
			ptr = ptr.next //iterate through the graph

		}
		if ptr.next != nil { //insert into the middle
			node.next = ptr.next.next
		}
		ptr.next = &node
	}
	return head
}
func main() {

	fmt.Println("New List Problems")
	head := CreateDoubleConnectedList(4)

	for i := 0; i < 10; i++ {
		head = InsertElement(head, i, 0)
	}

	fmt.Println("The list:")
	DisplayList(head)
	head = InsertElement(head, 44, 11)
	head = InsertElement(head, 33, 3)
	fmt.Println("\nList after insertions:")
	DisplayList(head)
	head = DeleteElement(head, 2)
	head = DeleteElement(head, 0)
	head = DeleteElement(head, 9)
	fmt.Println("\nList after Deletions:")
	DisplayList(head)
}
