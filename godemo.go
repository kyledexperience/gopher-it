package main

import "fmt"

func main() {

	messageLen := 10

	fmt.Println("Trying to send message of length: ", messageLen)

	if maxMessageLen := 20; maxMessageLen > messageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message too long")
	}

}
