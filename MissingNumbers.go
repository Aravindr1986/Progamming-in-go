package main

/*
Author @ Aravind
Find the missing number in a given integer array
*/
import (
	"fmt"
)

func main() {

	arr := []int{2, 3, 5, 5, 5, 8, 9, 10, 13, 15}

	var i, counter int = 0, 1
	fmt.Println("Missing Numbers:")
	for i = 0; i < len(arr); i++ {
		//flag := 0
		if arr[i] > counter { //check is counter value is missed in the array
			fmt.Print(" ")
			fmt.Print(counter) //print the counter as the missing value
			i--                //wait for the counter to catch up to the current number
			counter++          //increment the counter
		} else if counter <= arr[i] { // increment the counter
			counter++
		}
	}

}
