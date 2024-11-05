package main

import "fmt"

func concat(s1, s2 string) string {
	return s1 + s2
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func increment(x int) int {
	x = x + 1
	return x
}

func main() {

	messageLen := 10

	fmt.Println("Trying to send message of length: ", messageLen)

	if maxMessageLen := 20; maxMessageLen > messageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message too long")
	}

	fmt.Println(concat("Bob", " The Builder"))
	fmt.Println("2 + 3 = ", sum(2, 3))

	x := 1
	x = increment(x)
	fmt.Println(x)

}
