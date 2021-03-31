package main

import "fmt"

func set(m chan string) {
	m <- "ping1"
}

func main() {
	message := make(chan string)

	go set(message)

	go func() {
		message <- "ping"
	}()

	msg := <-message

	fmt.Println(msg)
}
