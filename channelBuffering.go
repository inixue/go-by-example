package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(message) // 地址

	fmt.Println(<-message)
	fmt.Println(<-message)
}
