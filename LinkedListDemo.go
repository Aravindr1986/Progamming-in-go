package main

import "github.com/Progamming-in-go/llist"
import "fmt"

func main() {

	fmt.Println("New List Problems")
	head := llist.CreateDoubleConnectedList(4)

	for i := 0; i < 10; i++ {
		head = llist.InsertElement(head, i, 0)
	}

	fmt.Println("The list:")
	llist.DisplayList(head)
	head = llist.InsertElement(head, 44, 11)
	head = llist.InsertElement(head, 33, 3)
	fmt.Println("\nList after insertions:")
	llist.DisplayList(head)
	head = llist.DeleteElement(head, 2)
	head = llist.DeleteElement(head, 0)
	head = llist.DeleteElement(head, 9)
	fmt.Println("\nList after Deletions:")
	llist.DisplayList(head)
}
