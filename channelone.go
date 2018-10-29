package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendData(msg chan string) {
	msg <- "hello routine one"
	wg.Done()
}
func receiveData(msg chan string) {
	str := <-msg
	fmt.Println("In second routnine")
	fmt.Println(str)
	wg.Done()
}
func main() {
	wg.Add(2)
	message := make(chan string)
	go sendData(message)
	//msg := <-message
	go receiveData(message)
	/*fmt.Println(msg)
	go func() {
		message <- "hello from routine Two"

	}()
	msg = <-message*/
	//fmt.Println(msg)
	wg.Wait()
	//fmt.Scanln()
}
