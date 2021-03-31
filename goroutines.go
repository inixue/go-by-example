package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i, time.Now())
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		now := time.Now()
		fmt.Println(msg, now.String())
	}("going")

	time.Sleep(time.Second) // 等待1s是为了等待 goroutine 执行并打印
	fmt.Println("done")

}
